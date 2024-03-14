package main

type Company struct {
	ID        int
	Name      string
	Parent    *Company
	Employees int
}

type ICompanyService interface {
	GetTopLevelParent(child *Company) *Company
	GetEmployeeCountForCompanyAndChildren(company *Company, allcompanies []*Company) int
}

type CompanyService struct{}

func (s *CompanyService) GetTopLevelParent(child *Company) *Company {
	if child == nil {
		return nil
	}

	for child.Parent != nil {
		child = child.Parent
	}
	return child
}

func (s *CompanyService) GetEmployeeCountForCompanyAndChildren(company *Company, companies []*Company) int {
	if company == nil {
		return 0
	}

	count := company.Employees

	for _, c := range companies {
		if c.Parent == company {
			count += s.GetEmployeeCountForCompanyAndChildren(c, companies)
		}
	}

	return count
}

var companies []*Company

func init() {
	company1 := &Company{ID: 1, Name: "Company 1", Employees: 3}
	company2 := &Company{ID: 2, Name: "Company 2", Employees: 3}
	company3 := &Company{ID: 3, Name: "Company 3", Employees: 2}
	company4 := &Company{ID: 4, Name: "Company 4", Employees: 3}
	company5 := &Company{ID: 5, Name: "Company 5", Employees: 1}
	company6 := &Company{ID: 6, Name: "Company 6", Employees: 2}
	company7 := &Company{ID: 7, Name: "Company 7", Employees: 5}
	company8 := &Company{ID: 8, Name: "Company 8", Employees: 4}
	company9 := &Company{ID: 9, Name: "Company 9", Employees: 2}
	company10 := &Company{ID: 10, Name: "Company 10", Employees: 3}

	company2.Parent = company1
	company3.Parent = company1
	company4.Parent = company2
	company5.Parent = company4
	company6.Parent = company3
	company7.Parent = company5
	company8.Parent = company5
	company9.Parent = company6
	company10.Parent = company6

	companies = append(companies, company1, company2, company3, company4, company5, company6, company7, company8, company9, company10)
}
