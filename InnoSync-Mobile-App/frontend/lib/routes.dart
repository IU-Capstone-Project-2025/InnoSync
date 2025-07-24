import 'screens/login_form.dart';
import 'screens/profile_creation_page.dart';
import 'screens/main_page.dart';
import 'screens/pages/settings_page.dart';

final appRoutes = {
  '/login': (context) => const LoginFormPage(),
  '/create_profile': (context) => const ProfileCreationPage(),
  '/main': (context) => const MainPage(),
  '/settings': (context) => const SettingsPage(),
};
