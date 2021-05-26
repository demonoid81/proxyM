<template>
  <i-form
    class="login-form"
    ref="loginForm"
    :model="form"
    :rules="rules"
    @keydown.enter.native="handleSubmit"
  >
    <i-form-item prop="username">
      <i-input v-model="form.username" placeholder="Please enter user name">
        <span slot="prefix">
          <i-icon :size="16" type="md-person" />
        </span>
      </i-input>
    </i-form-item>
    <i-form-item prop="password">
      <i-input type="password" v-model="form.password" placeholder="Please enter the password">
        <span slot="prefix">
          <i-icon :size="14" type="md-lock" />
        </span>
      </i-input>
    </i-form-item>
    <i-form-item>
      <i-button @click="handleSubmit" type="primary" long :loading="loading">Login</i-button>
    </i-form-item>
  </i-form>
</template>

<script>
import { getCaptchaidApi, getCaptchaBase64 } from '@/api/personal-center/user'

export default {
  name: 'LoginForm',

  props: {
    loading: {
      type: Boolean,
      default: false
    },
    userNameRules: {
      type: Array,
      default: () => {
        return [{ required: true, message: 'Account cannot be empty', trigger: 'blur' }]
      }
    },
    passwordRules: {
      type: Array,
      default: () => {
        return [{ required: true, message: 'Password can not be blank', trigger: 'blur' }]
      }
    },
    reloadCaptcha: String
  },

  data() {
    return {
      form: {
        username: 'admin',
        password: '123456',
      }
    }
  },

  computed: {
    rules() {
      return {
        username: this.userNameRules,
        password: this.passwordRules,
      }
    }
  },

  methods: {
    handleSubmit() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.$emit('on-submit', {
            username: this.form.username,
            password: this.form.password,
          })
        }
      })
    }
  }
}
</script>

<style lang="less">
.login-form {
  &-item {
    margin-bottom: 40px;

    &:last-child {
      margin-bottom: 10px;
    }
  }

  &__captcha {
    height: 33px;
    cursor: pointer;
  }
}
</style>
