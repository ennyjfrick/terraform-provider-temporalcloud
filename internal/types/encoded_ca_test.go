package types

import "testing"

func TestCACertNormalization(t *testing.T) {
	input := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJ5VENDQVZDZ0F3SUJBZ0lSQVdIa0MrNkpVZjNzOVRxNDNtZHAyemd3Q2dZSUtvWkl6ajBFQXdNd0V6RVIKTUE4R0ExVUVDaE1JZEdWdGNHOXlZV3d3SGhjTk1qTXdPREV3TURBd09UUTFXaGNOTWpRd09EQTVNREF4TURRMQpXakFUTVJFd0R3WURWUVFLRXdoMFpXMXdiM0poYkRCMk1CQUdCeXFHU000OUFnRUdCU3VCQkFBaUEySUFCQ3pRCjdEd3dHU1FLTTZacngzUXR3N0l1YmZ4aUozUlNYQ3FtY0doRWJGVmVvY3dBZEVnTVlsd1NsVWlXdERaVlIyZE0KWE05VVpMV0s0YUdHbkROUzVNaGN6NmliU0JTN093ZjR0UlpaQTlTcEZDak53MkhyYWFpVVZWK0VVZ3hvZTZObwpNR1l3RGdZRFZSMFBBUUgvQkFRREFnR0dNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGRzROCjhsSVhxUUt4d1ZzL2l4VnpkRjZYR1ptK01DUUdBMVVkRVFRZE1CdUNHV05zYVdWdWRDNXliMjkwTG5SbGJYQnYKY21Gc0xsQjFWSE13Q2dZSUtvWkl6ajBFQXdNRFp3QXdaQUl3UkxmbTlTN3JLR2QzMEtkUXZVTWNPY0RKbG1Edwo2L29NNlVPSkZ4TGVHY3BZYmd4US9iRml6ZStZeDlROWtOZU1BakE3R2lGc2FpcGFLdFdIeTVNQ09DYXMzWlA2Cit0dExhWE5Yc3MzWjVXazV2aERRbnlFOEpSM3JQZVEyY0hYTGlBMD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoKCi0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlCeGpDQ0FVMmdBd0lCQWdJUkE3b2E2dnhqd3RvVHdFdjNhVnZoZWh3d0NnWUlLb1pJemowRUF3TXdFakVRCk1BNEdBMVVFQ2hNSGRHVnpkR2x1WnpBZUZ3MHlOREE0TURJeE5qUXhOVFJhRncweU5UQTRNREl4TmpReU5UUmEKTUJJeEVEQU9CZ05WQkFvVEIzUmxjM1JwYm1jd2RqQVFCZ2NxaGtqT1BRSUJCZ1VyZ1FRQUlnTmlBQVFDcVVqVApEUVVKN2t3a055K3hkc0l3TitEY2hvSmJjdWVQVk9FQTB5STR0M05jS3lDcDJSTjhkbVAzbjFidVhtVVFNODBFCmxsQVlNaDFHcEU3VW1oT1l2aVVXenFWajNmN0s1Qm8wT1QvUjFxcndxVldXL0ZvbU5vdVlxK3o4TUxTalp6QmwKTUE0R0ExVWREd0VCL3dRRUF3SUJoakFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQjBHQTFVZERnUVdCQlJPVS9FSAp4Q1dDWmNJemVTcnR0NGhIV1ozY3VUQWpCZ05WSFJFRUhEQWFnaGhqYkdsbGJuUXVjbTl2ZEM1MFpYTjBhVzVuCkxsRnNSRk13Q2dZSUtvWkl6ajBFQXdNRFp3QXdaQUl3VlJJSEFzbmExajdUZnllQWR4YUNpY2dNK1lHcTQwVTQKUTdjOThCVlg3M1h1NkFnWXlBVU41eGlvbkJZSklCU3FBakFMcWRkanFzVG9pN0hXMTd5STVuN1VzSUdabHdrcQoxMm9rQjFJR0JSbU9pOU1SSnhuVk0wZXhrWGFUaEhGZ0toYz0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ=="
	expected := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJ5VENDQVZDZ0F3SUJBZ0lSQVdIa0MrNkpVZjNzOVRxNDNtZHAyemd3Q2dZSUtvWkl6ajBFQXdNd0V6RVIKTUE4R0ExVUVDaE1JZEdWdGNHOXlZV3d3SGhjTk1qTXdPREV3TURBd09UUTFXaGNOTWpRd09EQTVNREF4TURRMQpXakFUTVJFd0R3WURWUVFLRXdoMFpXMXdiM0poYkRCMk1CQUdCeXFHU000OUFnRUdCU3VCQkFBaUEySUFCQ3pRCjdEd3dHU1FLTTZacngzUXR3N0l1YmZ4aUozUlNYQ3FtY0doRWJGVmVvY3dBZEVnTVlsd1NsVWlXdERaVlIyZE0KWE05VVpMV0s0YUdHbkROUzVNaGN6NmliU0JTN093ZjR0UlpaQTlTcEZDak53MkhyYWFpVVZWK0VVZ3hvZTZObwpNR1l3RGdZRFZSMFBBUUgvQkFRREFnR0dNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGRzROCjhsSVhxUUt4d1ZzL2l4VnpkRjZYR1ptK01DUUdBMVVkRVFRZE1CdUNHV05zYVdWdWRDNXliMjkwTG5SbGJYQnYKY21Gc0xsQjFWSE13Q2dZSUtvWkl6ajBFQXdNRFp3QXdaQUl3UkxmbTlTN3JLR2QzMEtkUXZVTWNPY0RKbG1Edwo2L29NNlVPSkZ4TGVHY3BZYmd4US9iRml6ZStZeDlROWtOZU1BakE3R2lGc2FpcGFLdFdIeTVNQ09DYXMzWlA2Cit0dExhWE5Yc3MzWjVXazV2aERRbnlFOEpSM3JQZVEyY0hYTGlBMD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQotLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQnhqQ0NBVTJnQXdJQkFnSVJBN29hNnZ4and0b1R3RXYzYVZ2aGVod3dDZ1lJS29aSXpqMEVBd013RWpFUQpNQTRHQTFVRUNoTUhkR1Z6ZEdsdVp6QWVGdzB5TkRBNE1ESXhOalF4TlRSYUZ3MHlOVEE0TURJeE5qUXlOVFJhCk1CSXhFREFPQmdOVkJBb1RCM1JsYzNScGJtY3dkakFRQmdjcWhrak9QUUlCQmdVcmdRUUFJZ05pQUFRQ3FValQKRFFVSjdrd2tOeSt4ZHNJd04rRGNob0piY3VlUFZPRUEweUk0dDNOY0t5Q3AyUk44ZG1QM24xYnVYbVVRTTgwRQpsbEFZTWgxR3BFN1VtaE9ZdmlVV3pxVmozZjdLNUJvME9UL1IxcXJ3cVZXVy9Gb21Ob3VZcSt6OE1MU2paekJsCk1BNEdBMVVkRHdFQi93UUVBd0lCaGpBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUIwR0ExVWREZ1FXQkJST1UvRUgKeENXQ1pjSXplU3J0dDRoSFdaM2N1VEFqQmdOVkhSRUVIREFhZ2hoamJHbGxiblF1Y205dmRDNTBaWE4wYVc1bgpMbEZzUkZNd0NnWUlLb1pJemowRUF3TURad0F3WkFJd1ZSSUhBc25hMWo3VGZ5ZUFkeGFDaWNnTStZR3E0MFU0ClE3Yzk4QlZYNzNYdTZBZ1l5QVVONXhpb25CWUpJQlNxQWpBTHFkZGpxc1RvaTdIVzE3eUk1bjdVc0lHWmx3a3EKMTJva0IxSUdCUm1PaTlNUkp4blZNMGV4a1hhVGhIRmdLaGM9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
	normalized, err := normalizeCAString(input)
	if err != nil {
		t.Fatalf("failed to normalize cert: %v", err)
	}

	if normalized != expected {
		t.Fatalf("unexpected normalized cert: %s", normalized)
	}
}
