package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type QueryValidator struct {
    validationErrors []string
}

func NewValidator() *QueryValidator {
    return &QueryValidator{}
}

func (v *QueryValidator) ValidateInt(value string, fieldName string) int {
    if value == "" {
        v.validationErrors = append(v.validationErrors, 
            fmt.Sprintf("%s cannot be empty", fieldName))
        return 0
    }

    intValue, err := strconv.Atoi(value)
    if err != nil {
        v.validationErrors = append(v.validationErrors, 
            fmt.Sprintf("%s must be a valid integer", fieldName))
        return 0
    }

    return intValue
}

func (v *QueryValidator) ValidateProjectType(typeStr string) string {
    if !isProjectType(typeStr) {
        v.validationErrors = append(v.validationErrors, fmt.Sprintf("\"%s\" is not a valid project type", typeStr))
        return ""
    }
    return typeStr
}

func (v *QueryValidator) ValidateProjectTypes(typesStr string) []string {
    if typesStr == "" {
        return nil
    }

    types := strings.Split(typesStr, ",")
    var validTypes []string

    for _, t := range types {
        if isProjectType(t) {
            validTypes = append(validTypes, t)
        } else {
            v.validationErrors = append(v.validationErrors, 
                fmt.Sprintf("\"%s\" is not a valid project type", t))
        }
    }

    return validTypes
}

func (v *QueryValidator) ValidateGenres(genresStr string) []string {
    if genresStr == "" {
        return nil
    }
    genres := strings.Split(genresStr, ",")
    var codes []string
    for _, g := range genres {
        if c, ok := GetCategoryByLink(g); ok {
            codes = append(codes, c.CateCode)
        } else {
            v.validationErrors = append(v.validationErrors, fmt.Sprintf("\"%s\" is not a valid genre", g))
        }
    }
    return codes
}

func (v *QueryValidator) Errors() error {
    if len(v.validationErrors) == 0 {
        return nil
    }
    return fmt.Errorf(strings.Join(v.validationErrors, "\n"))
}

func isProjectType(t string) bool {
    switch t {
    case "manga": return true
    case "novel": return true
    case "comic": return true
    }
    return false
}
