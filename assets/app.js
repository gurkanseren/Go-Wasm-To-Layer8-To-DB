new Vue({
  el: "#app",
  data: {
    // Data properties (for future use)
  },
  methods: {
    showLoginForm() {
      this.showLoginForm = true;
      this.showInitialButtons = false;
      this.showRegisterForm = false;
    },
    showRegisterForm() {
      this.showRegisterForm = true;
      this.showInitialButtons = false;
      this.showLoginForm = false;
    },
    goBack() {
      this.showLoginForm = false;
      this.showRegisterForm = false;
      this.showInitialButtons = true;
      // Reset other form fields and messages
    },
    // Other methods (for future use)
  },
});
