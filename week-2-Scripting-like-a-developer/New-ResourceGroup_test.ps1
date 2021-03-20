#to run the test
#invoke-Pester .\New-ResourceGroup_test.ps1

Describe "New-ResourceGroup" {
    $location = 'eastus2'
    $name = 'demo0221'

    It "Name should be demo0221" {
        $name | Should Be 'demo0221'
    }

    It "Location should be eastus2" {
        $location | Should Be 'eastus2'
    }
}