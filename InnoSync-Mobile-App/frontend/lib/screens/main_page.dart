import 'package:flutter/material.dart';
import 'package:frontend/models/invitation_model.dart';
import 'pages/discover_page.dart';
import 'pages/dashboard_page.dart';
import 'pages/invitations_page.dart';
import 'pages/my_projects_page.dart';
import '../widgets/common/main_app_bar.dart';
import 'pages/views/invitation_view.dart';
import 'pages/views/project_view.dart';
import 'pages/views/talent_view.dart';
import '../models/project_model.dart';
import '../models/talent_model.dart';
import '../utils/token_storage.dart';

class MainPage extends StatefulWidget {
  const MainPage({super.key});

  @override
  State<MainPage> createState() => _MainPageState();
}

class _MainPageState extends State<MainPage> {
  int _currentIndex = 0;
  Project? _selectedProject;
  Talent? _selectedTalent;
  Invitation? _selectedInvitation;
  final List<Widget> _pages = [];

  @override
  void initState() {
    super.initState();
    getToken().then((token) {
      _pages.addAll([
        DiscoverPage(
          onProjectTap: (project) {
            setState(() {
              _selectedProject = project;
              _selectedTalent = null;
            });
          },
          onTalentTap: (talent) {
            setState(() {
              _selectedTalent = talent;
              _selectedProject = null;
            });
          },
        ),
        InvitationsPage(
          onInvitationTap: (invitation) {
            setState(() {
              _selectedInvitation = invitation;
            });
          },
        ),
        MyProjectsPage(
          onProjectTap: (project) {
            setState(() {
              _selectedProject = project;
            });
          },
        ),
        if (token != null) DashboardPage(token: token),
      ]);
      setState(() {});
    });
  }

  void _goBackFromView() {
    setState(() {
      _selectedProject = null;
      _selectedTalent = null;
      _selectedInvitation = null;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_pages.isEmpty) {
      return const Scaffold(body: Center(child: CircularProgressIndicator()));
    }
    return Scaffold(
      appBar: const MainAppBar(title: 'InnoSync'),
      body: _selectedProject != null
          ? ProjectView(project: _selectedProject!, onBack: _goBackFromView)
          : _selectedTalent != null
          ? TalentView(talent: _selectedTalent!, onBack: _goBackFromView)
          : _selectedInvitation != null
          ? InvitationView(invitation: _selectedInvitation!)
          : _pages[_currentIndex],
      bottomNavigationBar: (_selectedProject == null && _selectedTalent == null)
          ? BottomNavigationBar(
              currentIndex: _currentIndex,
              onTap: (index) => setState(() => _currentIndex = index),
              selectedItemColor: Theme.of(context).colorScheme.primary,
              unselectedItemColor: Theme.of(
                context,
              ).colorScheme.onSurface.withValues(alpha: 0.6),
              backgroundColor: Theme.of(context).colorScheme.surface,
              type: BottomNavigationBarType.fixed,
              elevation: 0,
              items: const [
                BottomNavigationBarItem(
                  icon: Icon(Icons.search),
                  label: 'Discover',
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.mail),
                  label: 'Invitations',
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.folder),
                  label: 'My Projects',
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.dashboard),
                  label: 'Dashboard',
                ),
              ],
            )
          : null,
    );
  }
}
