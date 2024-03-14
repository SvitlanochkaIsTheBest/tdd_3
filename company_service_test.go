package main

import (
	"reflect"
	"testing"
)

func TestCompanyService_GetTopLevelParent(t *testing.T) {
	tests := []struct {
		name  string
		child *Company
		want  *Company
	}{
		{
			name:  "Nil child",
			child: nil,
			want:  nil,
		},
		{
			name:  "Root company",
			child: companies[0],
			want:  companies[0],
		},
		{
			name:  "Child company",
			child: companies[8],
			want:  companies[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CompanyService{}
			if got := s.GetTopLevelParent(tt.child); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyService.GetTopLevelParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyService_GetEmployeeCountForCompanyAndChildren(t *testing.T) {
	tests := []struct {
		name      string
		company   *Company
		companies []*Company
		want      int
	}{
		{
			name:      "Nil company",
			company:   nil,
			companies: companies,
			want:      0,
		},
		{
			name:      "Company 1",
			company:   companies[0],
			companies: companies,
			want:      28,
		},
		{
			name:      "Company 2",
			company:   companies[1],
			companies: companies,
			want:      16,
		},
		{
			name:      "Company 3",
			company:   companies[2],
			companies: companies,
			want:      9,
		},
		{
			name:      "Company 4",
			company:   companies[3],
			companies: companies,
			want:      13,
		},
		{
			name:      "Company 5",
			company:   companies[4],
			companies: companies,
			want:      10,
		},
		{
			name:      "Company 6",
			company:   companies[5],
			companies: companies,
			want:      7,
		},
		{
			name:      "Company 7",
			company:   companies[6],
			companies: companies,
			want:      5,
		},
		{
			name:      "Company 8",
			company:   companies[7],
			companies: companies,
			want:      4,
		},
		{
			name:      "Company 9",
			company:   companies[8],
			companies: companies,
			want:      2,
		},
		{
			name:      "Company 10",
			company:   companies[9],
			companies: companies,
			want:      3,
		},
		{
			name:      "Non-existent company",
			company:   &Company{ID: 11, Name: "Company 11", Employees: 1},
			companies: companies,
			want:      1,
		},
		{
			name:      "Empty list of companies",
			company:   companies[0],
			companies: []*Company{},
			want:      3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CompanyService{}
			if got := s.GetEmployeeCountForCompanyAndChildren(tt.company, tt.companies); got != tt.want {
				t.Errorf("CompanyService.GetEmployeeCountForCompanyAndChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
