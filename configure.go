package potens

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
	"github.com/kubex/proto-go/discovery"
)

//SetIdentity Set your identity on the application
func (app *Application) SetIdentity(ident *identity.AppIdentity) error {
	if ident == nil {
		ident = &identity.AppIdentity{}
		err := ident.FromJSONFile(app.relPath("app-identity.json"))
		if err != nil {
			err = ident.FromJSONFile(app.relPath("_kubex/app-identity.json"))
			if err != nil {
				return err
			}
		}
	}

	block, _ := pem.Decode([]byte(ident.PrivateKey))
	if block == nil {
		return errors.New("No RSA private key found")
	}

	var key *rsa.PrivateKey
	if block.Type == "RSA PRIVATE KEY" {
		rsapk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return errors.New("Unable to read RSA private key")
		}
		key = rsapk
	}

	if !app.canBecomeGlobalAppID(ident.AppID) {
		return errors.New("The App ID specified in your identity does not match your definition")
	}

	app.identity = ident
	app.pk = key
	app.kh = ident.KeyHandle

	return nil
}

//SetDefinition Set your application definition
func (app *Application) SetDefinition(def *definition.AppDefinition) error {
	if def == nil {
		def = &definition.AppDefinition{}
		err := def.FromConfig(app.relPath("app-definition.yaml"))
		if err != nil {
			err = def.FromConfig(app.relPath("_kubex/app-definition.yaml"))
			if err != nil {
				return err
			}
		}
	}

	if len(def.VendorID) < 2 {
		return errors.New("The Vendor ID specified in your definition file is invalid")
	}

	if len(def.AppID) < 2 {
		return errors.New("The App ID specified in your definition file is invalid")
	}

	if !app.canBecomeGlobalAppID(def.GlobalAppID()) {
		return errors.New("The App ID specified in your definition does not match your identity")
	}

	switch def.Release {
	case definition.AppReleaseStable:
		app.appRelease = discovery.AppRelease_STABLE
		break
	case definition.AppReleaseBeta:
		app.appRelease = discovery.AppRelease_BETA
		break
	case definition.AppReleaseAlpha:
		app.appRelease = discovery.AppRelease_ALPHA
		break
	case definition.AppReleasePreAlpha:
		app.appRelease = discovery.AppRelease_PRE_ALPHA
		break
	}

	app.definition = def
	return nil
}

func (app *Application) canBecomeGlobalAppID(globalAppID string) bool {
	if app.identity != nil {
		return app.identity.AppID == globalAppID
	}

	if app.definition != nil {
		return app.definition.GlobalAppID() == globalAppID
	}

	return true
}

// Identity retrieves your identity
func (app *Application) Identity() *identity.AppIdentity {
	return app.identity
}

// Definition retrieves your definition
func (app *Application) Definition() *definition.AppDefinition {
	return app.definition
}
