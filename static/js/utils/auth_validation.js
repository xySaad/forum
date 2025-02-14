const AUTH_RULES = {
  USERNAME: {
    MIN_LENGTH: 3,
    MAX_LENGTH: 20,
    ALLOWED_CHARS: /^[a-zA-Z0-9.]+$/,
    ERRORS: {
      LENGTH: "Username must be between 3 and 20 characters",
      INVALID: "Username must contain only alphanumeric characters and dot (.)",
    },
  },
  PASSWORD: {
    MIN_LENGTH: 12,
    MAX_LENGTH: 256,
    ERRORS: {
      LENGTH: "Password must be between 12 and 256 characters",
      MISMATCH: "Passwords don't match",
    },
  },
  EMAIL: {
    PATTERN: /^[a-zA-Z0-9.]+@([a-zA-Z0-9]+\.)+[a-zA-Z0-9]{2,24}$/,
    ERRORS: {
      EMPTY: "Email can't be empty",
      INVALID: "Invalid email format",
    },
  },
};
const validateUsername = (username) => {
  if (
    username.length < AUTH_RULES.USERNAME.MIN_LENGTH ||
    username.length > AUTH_RULES.USERNAME.MAX_LENGTH
  ) {
    return AUTH_RULES.USERNAME.ERRORS.LENGTH;
  }
  if (!AUTH_RULES.USERNAME.ALLOWED_CHARS.test(username)) {
    return AUTH_RULES.USERNAME.ERRORS.INVALID;
  }
  return null;
};

const validateEmail = (email, isRegistration) => {
  if (!isRegistration) {
    return;
  }
  if (!email) return AUTH_RULES.EMAIL.ERRORS.EMPTY;
  if (!AUTH_RULES.EMAIL.PATTERN.test(email)) {
    return AUTH_RULES.EMAIL.ERRORS.INVALID;
  }
  return null;
};

const validatePassword = (password, confirmPassword, isRegistration) => {
  if (
    password.length < AUTH_RULES.PASSWORD.MIN_LENGTH ||
    password.length > AUTH_RULES.PASSWORD.MAX_LENGTH
  ) {
    return AUTH_RULES.PASSWORD.ERRORS.LENGTH;
  }
  if (isRegistration && password !== confirmPassword) {
    return AUTH_RULES.PASSWORD.ERRORS.MISMATCH;
  }
  return null;
};

export { validateUsername, validateEmail, validatePassword };
