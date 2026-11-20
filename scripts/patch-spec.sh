#!/bin/bash
# Patches api/openapi.json before generation. The chat completion stream chunk
# schema pins service_tier and usage to Go types (ServiceTier, CompletionUsage)
# that exist nowhere in the generated code. Point them at the schemas the
# non-streaming response already uses.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR
trap cleanup EXIT

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

readonly SCHEMA="OrchestratorChatCreateChatCompletionStreamResponse"
SPEC="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api/openapi.json"
readonly SPEC
TMP=""

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

cleanup() {
	[[ -z "$TMP" ]] || rm -f -- "$TMP"
}

main() {
	local tmp
	tmp="$(mktemp "${SPEC}.XXXXXX")"
	TMP="$tmp"

	jq --arg schema "$SCHEMA" '
		.components.schemas[$schema].properties |= (
			if has("service_tier") then
				.service_tier = {"$ref": "#/components/schemas/OrchestratorChatServiceTier"}
			else . end
			| if has("usage") then
				.usage = {
					"allOf": [{"$ref": "#/components/schemas/OrchestratorChatCompletionUsage"}],
					"description": .usage.description
				}
			else . end
		)
	' "$SPEC" >"$tmp"

	chmod 0644 "$tmp"
	mv -- "$tmp" "$SPEC"
}

main "$@"
