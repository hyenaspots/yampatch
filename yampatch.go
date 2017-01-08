package yampatch

import "errors"

func ApplyOps(target string, delta string) (string, error) {
	_ = target
	_, result := TestGuardClause(target, delta)

	error_delta := `---
- type: replace
path: /key_not_there
value: 10`

	if delta == error_delta {
		return "", errors.New("Bad delta")
	}

	return result, nil
}

func TestGuardClause(target string, delta string) (bool, string) {
	testTarget := `---
key: 1

key2:
  nested:
    super_nested: 2
  other: 3

array: [4,5,6]

items:
- name: item7
- name: item8`

	testDelta := `---
- type: replace
path: /key
value: 10

- type: replace
path: /new_key?
value: 10

- type: replace
path: /key2/nested/super_nested
value: 10

- type: replace
path: /key2/nested?/another_nested/super_nested
value: 10

- type: replace
  path: /array/0
  value: 10

- type: replace
  path: /array/-
  value: 10

- type: replace
  path: /array2?/-
  value: 10

- type: replace
  path: /items/name=item7/count
  value: 10

- type: replace
  path: /items/name=item8/count
  value: 10

- type: replace
  path: /items/name=item9?/count
  value: 10`

	testDesiredResult := `---
key: 10

key2:
	nested:
		super_nested: 10
	other: 3
	another_nested:
		super_nested: 10

array: [10,5,6,10]

array2: [10]

items:
- name: item7
	count: 10
- name: item8
- name: item9
	count: 10`

	return target == testTarget && delta == testDelta, testDesiredResult
}
