package identity

import (
	. "launchpad.net/gocheck"
	"launchpad.net/goose/testing/envsuite"
	"os"
)

type CredentialsTestSuite struct {
	// Isolate all of these tests from the real Environ.
	envsuite.EnvSuite
}

var _ = Suite(&CredentialsTestSuite{})

func (s *CredentialsTestSuite) TestCredentialsFromEnv(c *C) {
	var scenarios = []struct {
		summary   string
		env       map[string]string
		username  string
		password  string
		accessKey string
		secretKey string
		tenant    string
		region    string
		authURL   string
	}{
		{summary: "Old 'NOVA' style creds",
			env: map[string]string{
				"NOVA_USERNAME":   "test-user",
				"NOVA_PASSWORD":   "test-pass",
				"NOVA_API_KEY":    "test-access-key",
				"EC2_SECRET_KEYS": "test-secret-key",
				"NOVA_PROJECT_ID": "tenant-name",
				"NOVA_REGION":     "region",
			},
			username: "test-user",
			password: "test-pass",
			accessKey: "test-access-key",
			secretKey: "test-secret-key",
			tenant:   "tenant-name",
			region:   "region",
		},
		{summary: "New 'OS' style environment",
			env: map[string]string{
				"OS_USERNAME":    "test-user",
				"OS_PASSWORD":    "test-pass",
				"OS_ACCESS_KEY":  "test-access-key",
				"OS_SECRET_KEY":  "test-secret-key",
				"OS_TENANT_NAME": "tenant-name",
				"OS_REGION_NAME": "region",
			},
			username: "test-user",
			password: "test-pass",
			accessKey: "test-access-key",
			secretKey: "test-secret-key",
			tenant:   "tenant-name",
			region:   "region",
		},
	}
	for _, scenario := range scenarios {
		for key, value := range scenario.env {
			os.Setenv(key, value)
		}
		creds := CredentialsFromEnv()
		c.Check(creds.URL, Equals, scenario.authURL)
		c.Check(creds.User, Equals, scenario.username)
		c.Check(creds.Secrets, Equals, scenario.password)
		c.Check(creds.AccessKey, Equals, scenario.accessKey)
		c.Check(creds.SecretKey, Equals, scenario.secretKey)
		c.Check(creds.Region, Equals, scenario.region)
		c.Check(creds.TenantName, Equals, scenario.tenant)
	}
}

func (s *CredentialsTestSuite) TestCompleteCredentialsFromEnvValid(c *C) {
	env := map[string]string{
		"OS_AUTH_URL":    "http://auth",
		"OS_USERNAME":    "test-user",
		"OS_PASSWORD":    "test-pass",
		"OS_ACCESS_KEY":    "test-access-key",
		"OS_SECRET_KEY":    "test-secret-key",
		"OS_TENANT_NAME": "tenant-name",
		"OS_REGION_NAME": "region",
	}
	for key, value := range env {
		os.Setenv(key, value)
	}
	creds, err := CompleteCredentialsFromEnv()
	c.Assert(err, IsNil)
	c.Check(creds.URL, Equals, "http://auth")
	c.Check(creds.User, Equals, "test-user")
	c.Check(creds.Secrets, Equals, "test-pass")
	c.Check(creds.AccessKey, Equals, "test-access-key")
	c.Check(creds.SecretKey, Equals, "test-secret-key")
	c.Check(creds.Region, Equals, "region")
	c.Check(creds.TenantName, Equals, "tenant-name")
}

// An error is returned if not all required environment variables are set.
func (s *CredentialsTestSuite) TestCompleteCredentialsFromEnvInvalid(c *C) {
	env := map[string]string{
		"OS_AUTH_URL":    "http://auth",
		"OS_USERNAME":    "test-user",
		"OS_ACCESS_KEY":    "test-access-key",
		"OS_SECRET_KEY":    "test-secret-key",
		"OS_TENANT_NAME": "tenant-name",
		"OS_REGION_NAME": "region",
	}
	for key, value := range env {
		os.Setenv(key, value)
	}
	_, err := CompleteCredentialsFromEnv()
	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Matches, "required environment variable not set.*: Secrets")
}

func (s *CredentialsTestSuite) TestCompleteCredentialsFromEnvKeypair(c *C) {
	env := map[string]string{
		"OS_AUTH_URL":    "http://auth",
		"OS_USERNAME":    "test-user",
		"OS_PASSWORD":    "test-pass",
		"OS_ACCESS_KEY":    "test-access-key",
		"OS_SECRET_KEY":    "test-secret-key",
		"OS_TENANT_NAME": "tenant-name",
		"OS_REGION_NAME": "region",
	}
	for key, value := range env {
		os.Setenv(key, value)
	}
	creds, err := CompleteCredentialsFromEnv()
	c.Assert(err, IsNil)
	c.Check(creds.URL, Equals, "http://auth")
	c.Check(creds.AccessKey, Equals, "test-access-key")
	c.Check(creds.SecretKey, Equals, "test-secret-key")
	c.Check(creds.Region, Equals, "region")
	c.Check(creds.TenantName, Equals, "tenant-name")
}

func (s *CredentialsTestSuite) TestCompleteCredentialsFromEnvKeypairCompatibleEnvVars(c *C) {
	env := map[string]string{
		"OS_AUTH_URL":    "http://auth",
		"OS_USERNAME":    "test-user",
		"OS_PASSWORD":    "test-pass",
		"NOVA_API_KEY":    "test-access-key",
		"EC2_SECRET_KEYS":    "test-secret-key",
		"OS_TENANT_NAME": "tenant-name",
		"OS_REGION_NAME": "region",
	}
	for key, value := range env {
		os.Setenv(key, value)
	}
	creds, err := CompleteCredentialsFromEnv()
	c.Assert(err, IsNil)
	c.Check(creds.URL, Equals, "http://auth")
	c.Check(creds.AccessKey, Equals, "test-access-key")
	c.Check(creds.SecretKey, Equals, "test-secret-key")
	c.Check(creds.Region, Equals, "region")
	c.Check(creds.TenantName, Equals, "tenant-name")
}
