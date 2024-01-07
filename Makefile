#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += SPINNER
SPINNER_MK_SUMMARY := go-corelibs/spinner
SPINNER_MK_VERSION := v1.0.1

include CoreLibs.mk
