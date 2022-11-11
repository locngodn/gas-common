package zk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/samuel/go-zookeeper/zk"
)

type Zookeeper struct {
	*zk.Conn
}

func New(config *Config) (*Zookeeper, error) {
	node := []string{fmt.Sprintf("%s:%d", config.Host, config.Port)}
	client, _, err := zk.Connect(node, time.Second) //*10)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to Zookeeper")
	}
	return &Zookeeper{client}, nil
}

func (c *Zookeeper) GetStringValue(path string, defaultVal string) (string, error) {
	existed, _, err := c.Exists(path)
	if err != nil {
		return "", err
	}
	if existed {
		children, _, _, err := c.GetW(path)
		if err != nil {
			return defaultVal, err
		}
		return string(children), nil
	} else {
		if len(defaultVal) > 0 {
			return defaultVal, nil
		} else {
			return "", fmt.Errorf("%s must be set", path)
		}
	}
}

func (c *Zookeeper) GetIntValue(path string, defaultVal int) (int, error) {
	val := ""
	if defaultVal != 0 {
		val = strconv.Itoa(defaultVal)
	}

	val, err := c.GetStringValue(path, val)

	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (c *Zookeeper) GetInt64Value(path string, defaultVal int64) (int64, error) {
	val := ""
	if defaultVal != 0 {
		val = strconv.FormatInt(defaultVal, 10)
	}

	val, err := c.GetStringValue(path, val)

	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (c *Zookeeper) GetValue(path string, i interface{}) error {
	existed, _, err := c.Exists(path)
	if err != nil {
		return err
	}
	if existed {
		children, _, _, err := c.GetW(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(children, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Zookeeper) GetRequireValue(path string, i interface{}) error {
	existed, _, err := c.Exists(path)
	if err != nil {
		return err
	}
	if existed {
		children, _, _, err := c.GetW(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(children, i)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s must be set", path)
	}
	return nil
}

func (c *Zookeeper) GetOptionValue(path string, i interface{}) error {
	existed, _, err := c.Exists(path)
	if err != nil {
		return err
	}
	if existed {
		children, _, _, err := c.GetW(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(children, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Zookeeper) GetOptionStringValue(path string) (string, error) {
	existed, _, err := c.Exists(path)
	if err != nil {
		return "", err
	}
	if existed {
		children, _, _, err := c.GetW(path)
		if err != nil {
			return "", err
		}
		return string(children), nil
	} else {
		return "", nil
	}
}

func (c *Zookeeper) GetRequireStringValue(path string) (string, error) {
	existed, _, err := c.Exists(path)
	if err != nil {
		return "", err
	}
	if existed {
		children, _, _, err := c.GetW(path)
		if err != nil {
			return "", err
		}
		return string(children), nil
	} else {
		return "", fmt.Errorf("%s must be set", path)
	}
}

func (c *Zookeeper) GetRequireBoolValue(path string) (bool, error) {
	str, err := c.GetRequireStringValue(path)
	if err != nil {
		return false, err
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (c *Zookeeper) GetOptionBoolValue(path string) (bool, error) {
	str, err := c.GetOptionStringValue(path)
	if err != nil {
		return false, err
	}
	if len(str) == 0 {
		return false, nil
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (c *Zookeeper) GetRequireIntValue(path string) (int, error) {
	val, err := c.GetRequireStringValue(path)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (c *Zookeeper) GetOptionIntValue(path string) (int, error) {
	str, err := c.GetOptionStringValue(path)
	if err != nil {
		return 0, err
	}
	if len(str) == 0 {
		return 0, nil
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}
