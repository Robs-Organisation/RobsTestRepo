Name:           robscalc
Version:        1.0
Release:        1%{?dist}
Summary:        a calculator to learn something about gobuffalo

License:        All rights reserved
URL:            https://github.com/Robs-Organisation/RobsTestRepo
Source0:        %{name}-%{version}.tar.gz

Requires:       postgresql

%description
This is a description

%prep
mkdir -p %{buildroot}/%{_bindir}
mkdir -p %{buildroot}/%{_sysconfdir}/%{name}
mkdir -p %{buildroot}/lib/systemd/system
mkdir -p %{buildroot}/%{_datadir}/%{name}/template

cp /builds/RobsTest/Repo/calculator/rpm/assets/%{name}  %{buildroot}/%{_bindir}/%{name}
cp /builds/RobsTest/Repo/calculator/rpm/assets/config.json %{buildroot}/%{_sysconfdir}/%{name}
cp /builds/RobsTest/Repo/calculator/rpm/assets/override.conf %{buildroot}/%{_sysconfdir}/%{name}
cp /builds/RobsTest/Repo/calculator/rpm/assets/config.env %{buildroot}/%{_sysconfdir}/%{name}
cp /builds/RobsTest/Repo/calculator/rpm/assets/%{name}.service %{buildroot}/lib/systemd/system
cp -r /builds/RobsTest/Repo/calculator/rpm/assets/template %{buildroot}/%{_datadir}/%{name}/

%setup -c

%post
sed -i "s/SESSION_SECRET=""/\SESSION_SECRET="$(openssl rand -hex 30)"/" %{_sysconfdir}/%{name}/override.conf

%files
%{_bindir}/%{name}
%{_datadir}/%{name}/template
/lib/systemd/system/

%config(noreplace)
%{_sysconfdir}/%{name}/*

%clean
rm -rf $%{_builddir}
rm -rf %{buildroot}/%{_bindir}/%{name}
rm -rf %{buildroot}/%{_sysconfdir}/%{name}
rm -rf %{buildroot}/lib/systemd/system/
rm -rf $%{buildroot}/%{_datadir}/%{name}/