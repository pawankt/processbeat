# Don't try fancy stuff like debuginfo, which is useless on binary-only
# packages. Don't strip binary too
# Be sure buildpolicy set to do nothing
%define        __spec_install_post %{nil}
%define          debug_package %{nil}
%define        __os_install_post %{_dbpath}/brp-compress

Summary: Custom Processbeat to get status of defined process list
Name: processbeat
Version: 1.1.0
Release: Linux
License: Self
Group: System/Monitoring
Source: %{name}-%{version}-%{release}.tar.gz
Url: https://www.elastic.co/guide/en/beats/libbeat/master/community-beats.html
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root

%description
%{summary}

%prep
%setup -q -n %{name}-%{version}-%{release}

%build
# Empty section.

%install

## usr
## Add shell script 
%{__install} -d -m 755 %{buildroot}%{_bindir}/
mv processbeat.sh %{buildroot}%{_bindir}/processbeat
chmod +x %{buildroot}%{_bindir}/processbeat

## Add Processbeat binary to /usr/share
%{__install} -d -m 755 %{buildroot}/usr/share/%{name}/
%{__install} -d -m 755 %{buildroot}/usr/share/%{name}/bin
mv processbeat %{buildroot}/usr/share/%{name}/bin

## Add Processbeat modules
#%{__install} -d -m 755 %{buildroot}/usr/share/%{name}/module
#cp -rp module/* %{buildroot}/usr/share/%{name}/module/

## etc
%{__install} -d -m 755 %{buildroot}%{_sysconfdir}/%{name}/
%{__install} -m 644 processbeat.yml %{buildroot}%{_sysconfdir}/%{name}/

## var
%{__install} -d -m 755 %{buildroot}/var/lib/%{name}/

%files
%defattr(-,root,root)

%{_bindir}/%{name}
%dir /etc/%{name}/
%config(noreplace) /etc/%{name}/*
%dir /usr/share/%{name}
%doc /usr/share/%{name}/*
%dir /var/lib/%{name}

%changelog
* Wed Jun 12 2019 Pavan K Tambabathula 1.1.0
- Added custom processbeat.yml configuration

EOF
