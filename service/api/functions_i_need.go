





func ValidateTokenAndGetUser(r *http.Request, logger *logrus.Entry) (string, error) {
    // Extract the Authorization header
    authorizationHeader := strings.Split(r.Header.Get("Authorization"), " ")
    if len(authorizationHeader) != 2 {
        return "", errors.New("invalid authorization header format")
    }

    authorizationType, authorizationToken := authorizationHeader[0], authorizationHeader[1]
    if authorizationType != BEARER {
        return "", errors.New("not bearer authentication")
    }

    if !regex_uuid.MatchString(authorizationToken) {
        return "", errors.New("invalid token format")
    }

    username, err := extractUsernameFromToken(authorizationToken)
    if err != nil {
        return "", err
    }

    return username, nil
}