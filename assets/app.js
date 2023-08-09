new Vue({
  el: "#app",
  data: {
    showLoginForm: false,
    showRegisterForm: false,
    showInitialButtons: true,
    showImage: false,
    loginUsername: "",
    loginPassword: "",
    selectedChoice: "1",
    registerUsername: "",
    registerPassword: "",
    repeatPassword: "",
    imageSrc: "", // Set the image source here
    loginMessage: "",
    repeatPasswordError: "",
    showPictureOptions: false,
  },
  methods: {
    logUser() {
      const username = this.loginUsername;
      const password = this.loginPassword;
      const choice = this.selectedChoice;
      // Call your Go function to login the user
      const loginResult = loginUser(username, password, choice);
      if (loginResult === "success") {
        this.loginMessage = "Login successful";
      } else {
        this.loginMessage = "Login failed";
      }
    },
    regUser() {
      const username = this.registerUsername;
      const password = this.registerPassword;
      const repeatPassword = this.repeatPassword;
      const repeatPasswordError = document.getElementById(
        "repeatPasswordError"
      );

      if (password !== repeatPassword) {
        this.repeatPasswordError = "Passwords do not match";
        return;
      }

      // Call your Go function to register the user
      const registerResult = registerUser(username, password);
      if (registerResult === "success") {
        this.repeatPasswordError = "User registered";
      } else {
        this.repeatPasswordError = "User already exists";
      }
    },
    goBack() {
      this.showLoginForm = false;
      this.showRegisterForm = false;
      this.showInitialButtons = true;
      this.showImage = false;
      this.loginMessage = "";
      this.repeatPasswordError = "";
    },
    showLoginFormSection() {
      this.showLoginForm = true;
      this.showRegisterForm = false;
      this.showInitialButtons = false;
      this.showImage = false;
    },
    showRegisterFormSection() {
      this.showLoginForm = false;
      this.showRegisterForm = true;
      this.showInitialButtons = false;
      this.showImage = false;
    },
  },
});
