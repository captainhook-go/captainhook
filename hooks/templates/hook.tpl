#!/bin/sh

# installed by CaptainHook {{.VERSION}}

INTERACTIVE="{{.INTERACTION}} "

{{.RUN_PATH}}captainhook $INTERACTIVE--configuration={{.CONFIGURATION}} hook {{.HOOK_NAME}} "$@" <&0