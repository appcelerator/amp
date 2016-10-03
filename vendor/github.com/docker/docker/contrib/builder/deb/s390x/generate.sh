#!/bin/bash
set -e

# This file is used to auto-generate Dockerfiles for making debs via 'make deb'
#
# usage: ./generate.sh [versions]
#    ie: ./generate.sh
#        to update all Dockerfiles in this directory
#    or: ./generate.sh ubuntu-xenial
#        to only update ubuntu-xenial/Dockerfile
#    or: ./generate.sh ubuntu-newversion
#        to create a new folder and a Dockerfile within it

cd "$(dirname "$(readlink -f "$BASH_SOURCE")")"

versions=( "$@" )
if [ ${#versions[@]} -eq 0 ]; then
	versions=( */ )
fi
versions=( "${versions[@]%/}" )

for version in "${versions[@]}"; do
	echo "${versions[@]}"
	distro="${version%-*}"
	suite="${version##*-}"
	from="s390x/${distro}:${suite}"

	mkdir -p "$version"
	echo "$version -> FROM $from"
	cat > "$version/Dockerfile" <<-EOF
		#
		# THIS FILE IS AUTOGENERATED; SEE "contrib/builder/deb/s390x/generate.sh"!
		#

		FROM $from

	EOF

	extraBuildTags='pkcs11'
	runcBuildTags=

	# this list is sorted alphabetically; please keep it that way
	packages=(
		apparmor # for apparmor_parser for testing the profile
		bash-completion # for bash-completion debhelper integration
		btrfs-tools # for "btrfs/ioctl.h" (and "version.h" if possible)
		build-essential # "essential for building Debian packages"
		curl ca-certificates # for downloading Go
		debhelper # for easy ".deb" building
		dh-apparmor # for apparmor debhelper
		dh-systemd # for systemd debhelper integration
		git # for "git commit" info in "docker -v"
		libapparmor-dev # for "sys/apparmor.h"
		libdevmapper-dev # for "libdevmapper.h"
		libltdl-dev # for pkcs11 "ltdl.h"
		libseccomp-dev  # for "seccomp.h" & "libseccomp.so"
		libsqlite3-dev # for "sqlite3.h"
		pkg-config # for detecting things like libsystemd-journal dynamically
		libsystemd-dev
	)

	case "$suite" in
		# s390x needs libseccomp 2.3.1
		xenial)
			# Ubuntu Xenial has libseccomp 2.2.3
			runcBuildTags="apparmor selinux"
			;;
		*)
			extraBuildTags+=' seccomp'
			runcBuildTags="apparmor selinux seccomp"
			;;
	esac

	# update and install packages
	echo "RUN apt-get update && apt-get install -y ${packages[*]} --no-install-recommends && rm -rf /var/lib/apt/lists/*" >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	awk '$1 == "ENV" && $2 == "GO_VERSION" { print; exit }' ../../../../Dockerfile >> "$version/Dockerfile"
	echo 'RUN curl -fSL "https://storage.googleapis.com/golang/go${GO_VERSION}.linux-s390x.tar.gz" | tar xzC /usr/local' >> "$version/Dockerfile"
	echo 'ENV PATH $PATH:/usr/local/go/bin' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	echo 'ENV AUTO_GOPATH 1' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	# print build tags in alphabetical order
	buildTags=$( echo "apparmor selinux $extraBuildTags" | xargs -n1 | sort -n | tr '\n' ' ' | sed -e 's/[[:space:]]*$//' )

	echo "ENV DOCKER_BUILDTAGS $buildTags" >> "$version/Dockerfile"
	echo "ENV RUNC_BUILDTAGS $runcBuildTags" >> "$version/Dockerfile"
done
