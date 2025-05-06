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
  FIRSTNAME: {
    MIN_LENGTH: 3,
    MAX_LENGTH: 20,
    ALLOWED_CHARS: /^[a-zA-Z]+$/,
    ERRORS: {
      LENGTH: "FIRSTNAME must be between 3 and 20 characters",
      INVALID: "FIRSTNAME must contain only alphanumeric characters",
    },
  },
  LASTNAME: {
    MIN_LENGTH: 3,
    MAX_LENGTH: 20,
    ALLOWED_CHARS: /^[a-zA-Z]+$/,
    ERRORS: {
      LENGTH: "LASTNAME must be between 3 and 20 characters",
      INVALID: "LASTNAME must contain only alphanumeric characters",
    },
  },
  GENDER: {
    ALLOWED_GENDERS: ["male", "female", "other", "prefer not to say"],
    ERRORS: {
      INVALID: "this gender not allowed",
    },
  },
  AGE: {
    MIN_LENGTH: 18,
    MAX_LENGTH: 99,
    ALLOWED_CHARS: /^[0-9]+$/,
    ERRORS: {
      LENGTH: "AGE must be between 18 and 99",
      INVALID: "AGE must contain only NUMBERS",
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

const validateUsername = (username, context) => {
  if (context === "login") {
    return;
  }

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


const validateFirstname = (firstname, context) => {
 if (context === "login") {
  return;
}
  if (
    firstname.length < AUTH_RULES.FIRSTNAME.MIN_LENGTH ||
    firstname.length > AUTH_RULES.FIRSTNAME.MAX_LENGTH
  ) {
    return AUTH_RULES.FIRSTNAME.ERRORS.LENGTH;
  }
  if (!AUTH_RULES.FIRSTNAME.ALLOWED_CHARS.test(firstname)) {
    return AUTH_RULES.FIRSTNAME.ERRORS.INVALID;
  }
  return null;
};

const validateLastname = (lastname, context) => {
  if (context === "login") {
    return;
  }
  if (
    lastname.length < AUTH_RULES.LASTNAME.MIN_LENGTH ||
    lastname.length > AUTH_RULES.LASTNAME.MAX_LENGTH
  ) {
    return AUTH_RULES.LASTNAME.ERRORS.LENGTH;
  }
  if (!AUTH_RULES.LASTNAME.ALLOWED_CHARS.test(lastname)) {
    return AUTH_RULES.LASTNAME.ERRORS.INVALID;
  }
  return null;
};

const validateGender = (gender, context) => {
  if (context === "login") {
    return;
  }
  if (!AUTH_RULES.GENDER.ALLOWED_GENDERS.includes(gender)) {
    return AUTH_RULES.GENDER.ERRORS.INVALID;
  }
  return null;
};

const validateAge = (age, context) => {
  if (context === "login") {
    return;
  }
  const numericAge = Number(age);
  if (!AUTH_RULES.AGE.ALLOWED_CHARS.test(age)) {
    return AUTH_RULES.AGE.ERRORS.INVALID;
  }
  if (
    numericAge < AUTH_RULES.AGE.MIN_LENGTH ||
    numericAge > AUTH_RULES.AGE.MAX_LENGTH
  ) {
    return AUTH_RULES.AGE.ERRORS.LENGTH;
  }
  return null;
};

export {
  validateUsername,
  validateEmail,
  validatePassword,
  validateFirstname,
  validateLastname,
  validateGender,
  validateAge,
};
