package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

var (
	// ErrRSANilPrivKey is the error for trying to sign a JWT with a nil private key.
	ErrRSANilPrivKey = errors.New("jwt: RSA private key is nil")
	// ErrRSANilPubKey is the error for trying to verify a JWT with a nil public key.
	ErrRSANilPubKey = errors.New("jwt: RSA public key is nil")
)

type RSA struct {
	priv *rsa.PrivateKey
	pub  *rsa.PublicKey

	hash crypto.Hash
	opts *rsa.PSSOptions
	pool *pool
}

func NewRSA(sha Hash, priv *rsa.PrivateKey, pub *rsa.PublicKey) *RSA {
	hh := sha.hash()
	return &RSA{
		priv: priv,
		pub:  pub,
		hash: hh,
		pool: newPool(hh.New),
	}
}

func (r *RSA) Sign(payload []byte) ([]byte, error) {
	if r.priv == nil {
		return nil, ErrRSANilPrivKey
	}
	return r.sign(payload)
}

func (r *RSA) Size() int {
	pub := r.pub
	if pub == nil {
		pub = r.priv.Public().(*rsa.PublicKey)
	}
	return pub.Size()
}

func (r *RSA) String() string {
	if r.opts != nil {
		switch r.hash {
		case crypto.SHA256:
			return MethodPS256
		case crypto.SHA384:
			return MethodPS384
		case crypto.SHA512:
			return MethodPS512
		default:
			return ""
		}
	}
	switch r.hash {
	case crypto.SHA256:
		return MethodRS256
	case crypto.SHA384:
		return MethodRS384
	case crypto.SHA512:
		return MethodRS512
	default:
		return ""
	}
}

func (r *RSA) Verify(payload, sig []byte) (err error) {
	if r.pub == nil {
		return ErrRSANilPubKey
	}
	if sig, err = decodeToBytes(sig); err != nil {
		return err
	}
	if err = r.verify(payload, sig); err != nil {
		return err
	}
	return nil
}

func (r *RSA) WithPSS() *RSA {
	r.opts = &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       r.hash,
	}
	return r
}

func (r *RSA) sign(payload []byte) ([]byte, error) {
	sum, err := r.pool.sign(payload)
	if err != nil {
		return nil, err
	}
	if r.opts != nil {
		return rsa.SignPSS(rand.Reader, r.priv, r.hash, sum, r.opts)
	}
	return rsa.SignPKCS1v15(rand.Reader, r.priv, r.hash, sum)
}

func (r *RSA) verify(payload, sig []byte) error {
	sum, err := r.pool.sign(payload)
	if err != nil {
		return err
	}
	if r.opts != nil {
		return rsa.VerifyPSS(r.pub, r.hash, sum, sig, r.opts)
	}
	return rsa.VerifyPKCS1v15(r.pub, r.hash, sum, sig)
}
