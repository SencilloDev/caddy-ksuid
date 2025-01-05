# Caddy Ksuids

## Usage

```
route /api/v2/testing {
	ksuid
	header {
		x-request-id {ksuid.id}
	}
}
```
