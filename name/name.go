package name

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strings"
)

// Config ...
type Config struct {
	GivenNameStartsWith  string
	FamilyNameStartsWith string
}

func NewCryptoSeed() (int64, error) {
	var b [8]byte
	if _, err := crypto_rand.Read(b[:]); err != nil {
		return -1, err
	}

	seed := int64(binary.LittleEndian.Uint16(b[:]))
	return seed, nil
}

func DefaultRNG() (*rand.Rand, error) {
	seed, err := NewCryptoSeed()
	if err != nil {
		return nil, err
	}

	rng := rand.New(rand.NewSource(seed))
	return rng, nil
}

func Generate(rng *rand.Rand) string {
	return GenerateWithConfig(rng, nil)
}

func GenerateWithConfig(rng *rand.Rand, config *Config) string {

	sirname, err := randomSirname(rng, config)
	if err != nil {
		return "[Error]: " + err.Error()
	}

	givenName, err := randomGivenname(rng, config)
	if err != nil {
		return "[Error]: " + err.Error()
	}

	return givenName + " " + sirname
}

func randomSirname(rng *rand.Rand, config *Config) (string, error) {
	if config == nil || config.FamilyNameStartsWith == "" {
		return familyNames[rng.Intn(len(familyNames))], nil
	}

	ns := []string{}

	if config.FamilyNameStartsWith != "" {
		for _, n := range familyNames {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(config.FamilyNameStartsWith)) {
				ns = append(ns, n)
			}
		}
	}

	if len(ns) == 0 {
		return "", fmt.Errorf("no matching sirnames found")
	}

	return ns[rng.Intn(len(ns))], nil
}

func randomGivenname(rng *rand.Rand, config *Config) (string, error) {

	if config == nil || config.GivenNameStartsWith == "" {
		return givenNames[rng.Intn(len(givenNames))], nil
	}

	ns := []string{}

	if config.GivenNameStartsWith != "" {
		for _, n := range givenNames {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(config.GivenNameStartsWith)) {
				ns = append(ns, n)
			}
		}
	}

	if len(ns) == 0 {
		return "", fmt.Errorf("no matching givennames found")
	}

	return ns[rng.Intn(len(ns))], nil
}
