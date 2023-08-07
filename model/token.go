package model

func (t OauthToken) InScopes(scopes []string) bool {
	return InScopes(t.scopes, scopes)
}

func InScopes(has, wants []string) bool {
	hasScope := func(scope string, scopes []string) bool {
		for _, s := range scopes {
			if scope == s {
				return true
			}
		}
		return false
	}
	if wants == nil {
		return true
	}
	for _, sc := range wants {
		if !hasScope(sc, has) {
			return false
		}
	}
	return true
}
