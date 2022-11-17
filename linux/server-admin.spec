Name: server-admin
Version: %{binary}
Release:  el7
Summary: Server Admin

Group: System Environment/Daemons
License: Server Admin
Vendor: http://www.123.com
Autoreqprov : no
Source0: server-admin-linux_amd64-%{binary}.tar.gz

%description
Server Admin

%prep
%setup -q

%install
mkdir -p %{buildroot}/var/lib/server-admin
mv * %{buildroot}/var/lib/server-admin/

%pre
[ -f /etc/profile.d/bash-record.sh ] && rm -f /etc/profile.d/bash-record.sh
[ -d /var/lib/server-admin/static ] && rm -f /var/lib/server-admin/static
echo clear old bash and static file

%post
cp /var/lib/server-admin/server-admin.service /usr/lib/systemd/system/
cp /var/lib/server-admin/bash-record.sh /etc/profile.d/
chmod 755 /etc/profile.d/bash-record.sh
chmod 755 /var/lib/server-admin/server-admin
# dos2unix /etc/profile.d/bash-record.sh


%preun
systemctl stop server-admin

%clean

%postun
rm -rf /var/lib/server-admin/server*
rm -rf /var/lib/server-admin/static
rm -f /var/lib/server-admin/README.md
rm -f /etc/profile.d/bash-record.sh
rm -f /var/lib/server-admin/*.sh
rm -f /usr/lib/systemd/system/server-admin.service
echo finished uninstall.

%files
%defattr (-,-,-,0755)
/var/lib/server-admin/*