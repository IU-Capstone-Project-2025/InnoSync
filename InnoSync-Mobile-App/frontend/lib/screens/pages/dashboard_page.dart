import 'package:flutter/material.dart';
import 'package:frontend/repositories/profile_repository.dart';
import 'package:frontend/services/api_service.dart';
import 'package:frontend/screens/pages/views/profile_content.dart';
import 'package:frontend/screens/profile_creation_page.dart';
import 'package:frontend/widgets/common/spacing.dart';

class DashboardPage extends StatelessWidget {
  final String token;
  const DashboardPage({super.key, required this.token});

  @override
  Widget build(BuildContext context) {
    final apiService = ApiService();

    return Scaffold(
      backgroundColor: Theme.of(context).scaffoldBackgroundColor,
      body: FutureBuilder<Map<String, dynamic>>(
        future: apiService.fetchProfile(token),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          }
          if (snapshot.hasError || snapshot.data == null) {
            return const Center(child: Text('Failed to load profile data'));
          }
          return ProfileContent(profile: snapshot.data!);
        },
      ),
    );
  }
}

class _DashboardOverviewTab extends StatefulWidget {
  final String token;
  const _DashboardOverviewTab({required this.token});

  @override
  State<_DashboardOverviewTab> createState() => _DashboardOverviewTabState();
}

class _DashboardOverviewTabState extends State<_DashboardOverviewTab> {
  int _projectsCount = 0;
  int _invitationsCount = 0;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _fetchStats();
  }

  Future<void> _fetchStats() async {
    setState(() => _loading = true);
    try {
      final projects = await ApiService.getUserProjects(widget.token);
      final api = ApiService();
      final invitations = await api.fetchInvitations(widget.token);
      setState(() {
        _projectsCount = projects.length;
        _invitationsCount = invitations.length;
        _loading = false;
      });
    } catch (e) {
      setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    if (_loading) {
      return const Center(child: CircularProgressIndicator());
    }
    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'Welcome to your Dashboard!',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          ),
          const SizedBox(height: 16),
          Row(
            children: [
              Icon(
                Icons.account_circle,
                size: 40,
                color: theme.colorScheme.primary,
              ),
              const SizedBox(width: 12),
              Expanded(
                child: Text(
                  'Here you can track your projects, proposals, and invitations.',
                  style: TextStyle(fontSize: 16),
                ),
              ),
            ],
          ),
          const SizedBox(height: 24),
          Row(
            children: [
              Expanded(
                child: Card(
                  color: theme.colorScheme.primary.withValues(alpha: 0.1),
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Column(
                      children: [
                        Icon(
                          Icons.folder,
                          size: 32,
                          color: theme.colorScheme.primary,
                        ),
                        VSpace.small(),
                        Text('My Projects: $_projectsCount'),
                      ],
                    ),
                  ),
                ),
              ),
              Expanded(
                child: Card(
                  color: theme.colorScheme.primary.withValues(alpha: 0.1),
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Column(
                      children: [
                        Icon(
                          Icons.mail,
                          size: 32,
                          color: theme.colorScheme.primary,
                        ),
                        VSpace.small(),
                        Text('Invitations: $_invitationsCount'),
                      ],
                    ),
                  ),
                ),
              ),
            ],
          ),
          const SizedBox(height: 24),
          // Дополнительная информация
          Card(
            color: theme.colorScheme.primary.withValues(alpha: 0.07),
            child: Padding(
              padding: const EdgeInsets.all(16),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'Additional info:',
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                  VSpace.small(),
                  Text(
                    '• You can edit your profile, manage your projects and invitations.',
                  ),
                  VSpace.small(),
                  Text('• Stay tuned for more analytics and features!'),
                ],
              ),
            ),
          ),
          VSpace.medium(),
          Center(
            child: ElevatedButton.icon(
              icon: const Icon(Icons.edit),
              label: const Text('Edit Profile'),
              onPressed: () async {
                final profileRepository = ProfileRepository(
                  apiService: ApiService(),
                );
                final profile = await profileRepository.fetchProfile(
                  widget.token,
                );
                if (!mounted) return;
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) =>
                        ProfileCreationPage(profileData: profile),
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}
