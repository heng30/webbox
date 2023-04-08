#!/bin/bash

SUBDIRS := ./frontend ./backend

RECURSIVE_DOWNLOAD_DEPENDENCE := @for subdir in $(SUBDIRS); \
				   do \
				   ( cd $$subdir && make download-dependence ) || exit 1; \
				   done

RECURSIVE_MAKE := @for subdir in $(SUBDIRS); \
				  do \
				  make -C $$subdir || exit 1; \
				  done

RECURSIVE_CLEAN := @for subdir in $(SUBDIRS); \
				   do \
				   ( cd $$subdir && make clean ) || exit 1; \
				   done

all:
	$(RECURSIVE_DOWNLOAD_DEPENDENCE)
	$(RECURSIVE_MAKE)

clean:
	$(RECURSIVE_CLEAN)

package:
	rm -rf webbox.tar.gz
	mkdir webbox
	cp -rf ./frontend/dist ./webbox/frontend
	cp -rf ./backend/webbox_* ./webbox
	tar -zcvf webbox.tar.gz ./webbox
	rm -rf webbox
