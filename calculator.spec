Name:           calculator
Version:        1.0
Release:        1%{?dist}
Summary:        Calculaor with Buffalo

License:        All rights reserved
Source0:        calculator-1.0.tar.gz

Requires:       postgresql

%description
This is a test to build an RPM with my calculator.

%prep

%setup -q

%install
install -m 0755 %{name} %{buildroot}%{_bindir}/%{name}

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root,-)
/CalcBuffaloBinary

%changelog
