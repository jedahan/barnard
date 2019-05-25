// Code generated by "enumer -type=Key -trimprefix=Key -yaml -json -transform=snake"; DO NOT EDIT.

//
package uiterm

import (
	"encoding/json"
	"fmt"
)

const (
	_KeyName_0 = "ctrl_tildectrl_actrl_bctrl_cctrl_dctrl_ectrl_fctrl_gbackspacetabctrl_jctrl_kctrl_lenterctrl_nctrl_octrl_pctrl_qctrl_rctrl_sctrl_tctrl_uctrl_vctrl_wctrl_xctrl_yctrl_zescctrl4ctrl5ctrl6ctrl7space"
	_KeyName_1 = "backspace2"
	_KeyName_2 = "mouse_rightmouse_middlemouse_leftarrow_rightarrow_leftarrow_downarrow_uppgdnpgupendhomedeleteinsertf12f11f10f9f8f7f6f5f4f3f2f1"
)

var (
	_KeyIndex_0 = [...]uint8{0, 10, 16, 22, 28, 34, 40, 46, 52, 61, 64, 70, 76, 82, 87, 93, 99, 105, 111, 117, 123, 129, 135, 141, 147, 153, 159, 165, 168, 173, 178, 183, 188, 193}
	_KeyIndex_1 = [...]uint8{0, 10}
	_KeyIndex_2 = [...]uint8{0, 11, 23, 33, 44, 54, 64, 72, 76, 80, 83, 87, 93, 99, 102, 105, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126}
)

func (i Key) String() string {
	switch {
	case 0 <= i && i <= 32:
		return _KeyName_0[_KeyIndex_0[i]:_KeyIndex_0[i+1]]
	case i == 127:
		return _KeyName_1
	case 65511 <= i && i <= 65535:
		i -= 65511
		return _KeyName_2[_KeyIndex_2[i]:_KeyIndex_2[i+1]]
	default:
		return fmt.Sprintf("Key(%d)", i)
	}
}

var _KeyValues = []Key{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 127, 65511, 65512, 65513, 65514, 65515, 65516, 65517, 65518, 65519, 65520, 65521, 65522, 65523, 65524, 65525, 65526, 65527, 65528, 65529, 65530, 65531, 65532, 65533, 65534, 65535}

var _KeyNameToValueMap = map[string]Key{
	_KeyName_0[0:10]:    0,
	_KeyName_0[10:16]:   1,
	_KeyName_0[16:22]:   2,
	_KeyName_0[22:28]:   3,
	_KeyName_0[28:34]:   4,
	_KeyName_0[34:40]:   5,
	_KeyName_0[40:46]:   6,
	_KeyName_0[46:52]:   7,
	_KeyName_0[52:61]:   8,
	_KeyName_0[61:64]:   9,
	_KeyName_0[64:70]:   10,
	_KeyName_0[70:76]:   11,
	_KeyName_0[76:82]:   12,
	_KeyName_0[82:87]:   13,
	_KeyName_0[87:93]:   14,
	_KeyName_0[93:99]:   15,
	_KeyName_0[99:105]:  16,
	_KeyName_0[105:111]: 17,
	_KeyName_0[111:117]: 18,
	_KeyName_0[117:123]: 19,
	_KeyName_0[123:129]: 20,
	_KeyName_0[129:135]: 21,
	_KeyName_0[135:141]: 22,
	_KeyName_0[141:147]: 23,
	_KeyName_0[147:153]: 24,
	_KeyName_0[153:159]: 25,
	_KeyName_0[159:165]: 26,
	_KeyName_0[165:168]: 27,
	_KeyName_0[168:173]: 28,
	_KeyName_0[173:178]: 29,
	_KeyName_0[178:183]: 30,
	_KeyName_0[183:188]: 31,
	_KeyName_0[188:193]: 32,
	_KeyName_1[0:10]:    127,
	_KeyName_2[0:11]:    65511,
	_KeyName_2[11:23]:   65512,
	_KeyName_2[23:33]:   65513,
	_KeyName_2[33:44]:   65514,
	_KeyName_2[44:54]:   65515,
	_KeyName_2[54:64]:   65516,
	_KeyName_2[64:72]:   65517,
	_KeyName_2[72:76]:   65518,
	_KeyName_2[76:80]:   65519,
	_KeyName_2[80:83]:   65520,
	_KeyName_2[83:87]:   65521,
	_KeyName_2[87:93]:   65522,
	_KeyName_2[93:99]:   65523,
	_KeyName_2[99:102]:  65524,
	_KeyName_2[102:105]: 65525,
	_KeyName_2[105:108]: 65526,
	_KeyName_2[108:110]: 65527,
	_KeyName_2[110:112]: 65528,
	_KeyName_2[112:114]: 65529,
	_KeyName_2[114:116]: 65530,
	_KeyName_2[116:118]: 65531,
	_KeyName_2[118:120]: 65532,
	_KeyName_2[120:122]: 65533,
	_KeyName_2[122:124]: 65534,
	_KeyName_2[124:126]: 65535,
}

// KeyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func KeyString(s string) (Key, error) {
	if val, ok := _KeyNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Key values", s)
}

// KeyValues returns all values of the enum
func KeyValues() []Key {
	return _KeyValues
}

// IsAKey returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Key) IsAKey() bool {
	for _, v := range _KeyValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Key
func (i Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Key
func (i *Key) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Key should be a string, got %s", data)
	}

	var err error
	*i, err = KeyString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Key
func (i Key) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Key
func (i *Key) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = KeyString(s)
	return err
}
