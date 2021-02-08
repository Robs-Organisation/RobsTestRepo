Name:           calculator
Version:        1.0
Release:        1%{?dist}
Summary:        Calculaor with Buffalo

License:        All rights reserved
Source0:        CalcBuffaloBinary-1.0.tar.gz

Requires:       postgresql

%description
This is a test to build an RPM with my calculator.

%prep
mkdir -p %{buildroot}%{_bindir}
mkdir -p %{buildroot}%{_sysconfdir}/%{name}

%setup -c

%install
install -d $RPM_BUILD_ROOT/opt/calculator
install home/runner/work/RobsCalculatorProject/RobsCalculatorProject/rpmbuild/sourceCode/CalcBuffaloBinary $RPM_BUILD_ROOT/opt/calculator/CalcBuffaloBinary

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root,-)
/opt/calculator/CalcBuffaloBinary

%changelog
