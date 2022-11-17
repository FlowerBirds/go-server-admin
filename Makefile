.PHONY: build

VERSION=0.0.9
BINARY=server-admin
PATH:=D:\06_devptools\go1.17.12.windows-amd64\go\bin:$(PATH)
GO111MODULE=on

# $(info $(PATH))
build_cmd = GOOS=$(1) GOARCH=$(2) CGO_ENABLED=1 CGO_LDFLAGS="-static" go build -ldflags "-X main.version=${VERSION}" -o build/${BINARY}-${VERSION}/$(BINARY)$(3)
tar = cd build && tar -cvzf ${BINARY}-$(1)_$(2)-${VERSION}.tar.gz ${BINARY}-${VERSION} && rm -rf ${BINARY}-${VERSION} && cd ..
cp_static = cp -r static build/${BINARY}-${VERSION}/ && cp -r README.md build/${BINARY}-${VERSION}/
cp_linux = cd linux && cp * ../build/${BINARY}-${VERSION}/
test = cd build && tar -xvzf ${BINARY}-$(1)_$(2)-${VERSION}.tar.gz


build: build/linux_amd64 build/rpm

build/linux_amd64:
	$(call build_cmd,linux,amd64,)
	$(call cp_static)
	$(call cp_linux)
	$(call tar,linux,amd64)


build/linux_arm64:
	$(call build_cmd,linux,arm64,)
	$(call cp_static)
	$(call cp_linux)
	$(call tar,linux,arm64)

build/windows_amd64:
	$(call build_cmd,windows,amd64,.exe)
	$(call cp_static)
	$(call tar,windows,amd64,.exe)

test/windows:
	$(call test,windows,amd64,.exe)

clean:
	cd build && rm -rf *

build/rpm:
	mkdir build/BUILD
	mkdir build/RPMS
	mkdir build/SOURCES
	mkdir build/SPECS
	mkdir build/SRPMS
	mkdir build/BUILDROOT
	cp build/${BINARY}-linux_amd64-${VERSION}.tar.gz build/SOURCES/
	cp linux/server-admin.spec build/SPECS/
	rpmbuild -ba --target=x86_64 build/SPECS/server-admin.spec --define "_topdir `pwd`/build/" --define "binary ${VERSION}"
	cp build/RPMS/x86_64/*.rpm build/
	rm -rf build/BUILD
	rm -rf build/RPMS
	rm -rf build/SOURCES
	rm -rf build/SPECS
	rm -rf build/SRPMS
	rm -rf build/BUILDROOT
