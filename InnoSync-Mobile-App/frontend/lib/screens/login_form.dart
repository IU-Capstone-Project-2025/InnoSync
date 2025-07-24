import 'package:flutter/material.dart';
import 'package:frontend/theme/dimensions.dart';
import 'package:frontend/widgets/common/spacing.dart';
import 'package:frontend/widgets/common/widgets.dart';
import 'package:frontend/widgets/login/widgets.dart';
import 'package:frontend/utils/ui_helpers.dart';
import 'package:frontend/services/api_service.dart';
import 'package:frontend/utils/token_storage.dart';
import 'package:frontend/widgets/common/theme_switcher_button.dart';

class LoginFormPage extends StatefulWidget {
  const LoginFormPage({super.key});

  @override
  State<LoginFormPage> createState() => _LoginFormPageState();
}

class _LoginFormPageState extends State<LoginFormPage>
    with SingleTickerProviderStateMixin {
  late TabController _tabController;
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  bool _rememberMe = false;
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final _fullNameController = TextEditingController();
  bool _obscurePassword = true;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 2, vsync: this);
  }

  @override
  void dispose() {
    _tabController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  void _logIn() async {
    final email = _emailController.text.trim();
    final password = _passwordController.text;
    try {
      final String? token = await ApiService.login(email, password);
      if (token != null) {
        await saveToken(token); // сохраняем токен для будущих запросов
        if (!mounted) return;
        Navigator.pushNamedAndRemoveUntil(context, '/main', (route) => false);
      } else {
        if (!mounted) return;
        UIHelpers.showError(
          context,
          'Login failed. Check your email and password.',
        );
      }
    } catch (e) {
      if (!mounted) return;
      UIHelpers.showError(context, 'Connection error.');
    }
  }

  void _signUp() async {
    final email = _emailController.text.trim();
    final password = _passwordController.text;
    final fullName = _fullNameController.text.trim();
    try {
      final success = await ApiService.signup(email, password, fullName);
      if (success) {
        // Immediately log in after signup
        final String? token = await ApiService.login(email, password);
        if (token != null) {
          await saveToken(token);
          if (!mounted) return;
          Navigator.pushNamed(
            context,
            '/create_profile',
            arguments: {'email': email, 'name': fullName},
          );
        } else {
          if (!mounted) return;
          UIHelpers.showError(context, 'Could not log in after signup.');
        }
      } else {
        if (!mounted) return;
        UIHelpers.showError(
          context,
          'Registration error. Email may already exist.',
        );
      }
    } catch (e) {
      if (!mounted) return;
      UIHelpers.showError(context, 'Server connection error.');
    }
  }

  Widget _buildLoginTab() {
    return LoginTab(
      emailController: _emailController,
      passwordController: _passwordController,
      rememberMe: _rememberMe,
      onRememberMeChanged: (value) =>
          setState(() => _rememberMe = value ?? false),
      onForgotPassword: () {
        // TODO: Implement forgot password
      },
      onLogin: _logIn,
    );
  }

  Widget _buildSignUpTab() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        ValidatedTextField(
          icon: Icons.person,
          hint: "Name",
          controller: _fullNameController,
          validator: (value) {
            if (value == null || value.trim().isEmpty) {
              return 'Enter Name';
            }
            return null;
          },
        ),
        const VSpace.medium(),
        ValidatedTextField(
          icon: Icons.email,
          hint: "Email",
          controller: _emailController,
          validator: (value) =>
              value == null || value.trim().isEmpty ? 'Enter Email' : null,
        ),
        const VSpace.medium(),
        ValidatedTextField(
          icon: Icons.lock,
          hint: "Password",
          obscure: _obscurePassword,
          controller: _passwordController,
          validator: (value) =>
              value == null || value.trim().isEmpty ? 'Enter Password' : null,
          toggleObscure: () {
            setState(() {
              _obscurePassword = !_obscurePassword;
            });
          },
        ),
        const VSpace.mediumPlus(),
        SubmitButton(
          text: "Continue",
          onPressed: () => _signUp(),
          isLoading: false,
        ),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).scaffoldBackgroundColor,
      body: FormContainer(
        formKey: _formKey,
        title: "InnoSync",
        subtitle: "We work it together !",
        child: Column(
          children: [
            TabSelector(
              controller: _tabController,
              tabLabels: const ['Log In', 'Sign Up'],
            ),
            const VSpace.xxl(),
            FixedHeightSpace(
              height: AppDimensions.tabViewHeight,
              child: TabBarView(
                controller: _tabController,
                children: [_buildLoginTab(), _buildSignUpTab()],
              ),
            ),
          ],
        ),
      ),
      floatingActionButton: const ThemeSwitcherButton(),
    );
  }
}
