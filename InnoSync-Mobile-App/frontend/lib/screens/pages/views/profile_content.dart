import 'package:flutter/material.dart';
import 'package:frontend/widgets/common/spacing.dart';

class ProfileContent extends StatelessWidget {
  final Map<String, dynamic> profile;
  const ProfileContent({super.key, required this.profile});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return SingleChildScrollView(
      padding: const EdgeInsets.all(24),
      child: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 800),
          child: Column(
            children: [
              // Profile Header Card
              Card(
                elevation: 8,
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(20),
                ),
                child: Container(
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(20),
                    gradient: LinearGradient(
                      begin: Alignment.topLeft,
                      end: Alignment.bottomRight,
                      colors: [
                        theme.colorScheme.primary.withOpacity(0.1),
                        theme.colorScheme.primary.withOpacity(0.05),
                      ],
                    ),
                  ),
                  padding: const EdgeInsets.all(32),
                  child: Row(
                    children: [
                      // Profile Avatar
                      Container(
                        decoration: BoxDecoration(
                          shape: BoxShape.circle,
                          border: Border.all(
                            color: theme.colorScheme.primary,
                            width: 4,
                          ),
                          boxShadow: [
                            BoxShadow(
                              color: theme.colorScheme.primary.withOpacity(0.3),
                              blurRadius: 12,
                              offset: const Offset(0, 4),
                            ),
                          ],
                        ),
                        child: CircleAvatar(
                          radius: 60,
                          backgroundImage: profile['avatar_url'] != null
                              ? NetworkImage(profile['avatar_url'])
                              : null,
                          child: profile['avatar_url'] == null
                              ? Icon(
                                  Icons.person,
                                  size: 60,
                                  color: theme.colorScheme.primary,
                                )
                              : null,
                        ),
                      ),
                      const HSpace.large(),
                      // Profile Info
                      Expanded(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              profile['name'] ?? 'No Name',
                              style: theme.textTheme.headlineMedium?.copyWith(
                                fontWeight: FontWeight.bold,
                                color: theme.colorScheme.onSurface,
                              ),
                            ),
                            const VSpace.small(),
                            Container(
                              padding: const EdgeInsets.symmetric(
                                horizontal: 16,
                                vertical: 8,
                              ),
                              decoration: BoxDecoration(
                                color: theme.colorScheme.primary,
                                borderRadius: BorderRadius.circular(20),
                              ),
                              child: Text(
                                profile['position'] ?? 'No Position',
                                style: theme.textTheme.titleMedium?.copyWith(
                                  color: theme.colorScheme.onPrimary,
                                  fontWeight: FontWeight.w600,
                                ),
                              ),
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
              ),

              const VSpace.large(),

              // Bio Section
              if (profile['bio'] != null &&
                  profile['bio'].toString().isNotEmpty) ...[
                _buildSectionCard(
                  theme,
                  'About Me',
                  Icons.info_outline,
                  Text(profile['bio'], style: theme.textTheme.bodyLarge),
                ),
                const VSpace.medium(),
              ],

              // Contact Information
              if (_hasContactInfo(profile)) ...[
                _buildSectionCard(
                  theme,
                  'Contact Information',
                  Icons.contact_mail,
                  Column(
                    children: [
                      if (profile['email'] != null)
                        _buildContactRow(
                          Icons.email,
                          'Email',
                          profile['email'],
                        ),
                      if (profile['github'] != null)
                        _buildContactRow(
                          Icons.code,
                          'GitHub',
                          profile['github'],
                        ),
                      if (profile['telegram'] != null)
                        _buildContactRow(
                          Icons.telegram,
                          'Telegram',
                          profile['telegram'],
                        ),
                    ],
                  ),
                ),
                const VSpace.medium(),
              ],

              // Technologies
              if (profile['technologies'] != null &&
                  profile['technologies'] is List &&
                  (profile['technologies'] as List).isNotEmpty) ...[
                _buildSectionCard(
                  theme,
                  'Technologies',
                  Icons.psychology,
                  Wrap(
                    spacing: 12,
                    runSpacing: 8,
                    children: (profile['technologies'] as List)
                        .map(
                          (tech) => Container(
                            padding: const EdgeInsets.symmetric(
                              horizontal: 16,
                              vertical: 8,
                            ),
                            decoration: BoxDecoration(
                              color: theme.colorScheme.primary.withOpacity(0.1),
                              borderRadius: BorderRadius.circular(20),
                              border: Border.all(
                                color: theme.colorScheme.primary.withOpacity(
                                  0.3,
                                ),
                              ),
                            ),
                            child: Text(
                              tech['name'] ?? tech.toString(),
                              style: theme.textTheme.bodyMedium?.copyWith(
                                color: theme.colorScheme.primary,
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                          ),
                        )
                        .toList(),
                  ),
                ),
                const VSpace.medium(),
              ],

              // Work Experience
              if (profile['work_experiences'] != null &&
                  profile['work_experiences'] is List &&
                  (profile['work_experiences'] as List).isNotEmpty) ...[
                _buildSectionCard(
                  theme,
                  'Work Experience',
                  Icons.work,
                  Column(
                    children: (profile['work_experiences'] as List)
                        .map(
                          (experience) => Container(
                            margin: const EdgeInsets.only(bottom: 16),
                            padding: const EdgeInsets.all(20),
                            decoration: BoxDecoration(
                              color: theme.colorScheme.surface,
                              borderRadius: BorderRadius.circular(16),
                              border: Border.all(
                                color: theme.colorScheme.outline.withOpacity(
                                  0.2,
                                ),
                              ),
                            ),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Row(
                                  children: [
                                    Icon(
                                      Icons.business,
                                      color: theme.colorScheme.primary,
                                      size: 20,
                                    ),
                                    const HSpace.small(),
                                    Expanded(
                                      child: Text(
                                        experience['position'] ??
                                            'Unknown Position',
                                        style: theme.textTheme.titleMedium
                                            ?.copyWith(
                                              fontWeight: FontWeight.bold,
                                              color:
                                                  theme.colorScheme.onSurface,
                                            ),
                                      ),
                                    ),
                                  ],
                                ),
                                const VSpace.small(),
                                Text(
                                  experience['company'] ?? 'Unknown Company',
                                  style: theme.textTheme.bodyLarge?.copyWith(
                                    color: theme.colorScheme.primary,
                                    fontWeight: FontWeight.w600,
                                  ),
                                ),
                                if (experience['start_date'] != null &&
                                    experience['end_date'] != null) ...[
                                  const VSpace.small(),
                                  Row(
                                    children: [
                                      Icon(
                                        Icons.calendar_today,
                                        size: 16,
                                        color: theme.colorScheme.onSurface
                                            .withOpacity(0.6),
                                      ),
                                      const HSpace.small(),
                                      Text(
                                        '${experience['start_date']} - ${experience['end_date']}',
                                        style: theme.textTheme.bodyMedium
                                            ?.copyWith(
                                              color: theme.colorScheme.onSurface
                                                  .withOpacity(0.6),
                                            ),
                                      ),
                                    ],
                                  ),
                                ],
                                if (experience['description'] != null) ...[
                                  const VSpace.medium(),
                                  Text(
                                    experience['description'],
                                    style: theme.textTheme.bodyMedium,
                                  ),
                                ],
                              ],
                            ),
                          ),
                        )
                        .toList(),
                  ),
                ),
              ],
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildSectionCard(
    ThemeData theme,
    String title,
    IconData icon,
    Widget content,
  ) {
    return Card(
      elevation: 4,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      child: Padding(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Container(
                  padding: const EdgeInsets.all(8),
                  decoration: BoxDecoration(
                    color: theme.colorScheme.primary.withOpacity(0.1),
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Icon(icon, color: theme.colorScheme.primary, size: 24),
                ),
                const HSpace.medium(),
                Text(
                  title,
                  style: theme.textTheme.headlineSmall?.copyWith(
                    fontWeight: FontWeight.bold,
                    color: theme.colorScheme.onSurface,
                  ),
                ),
              ],
            ),
            const VSpace.medium(),
            content,
          ],
        ),
      ),
    );
  }

  Widget _buildContactRow(IconData icon, String label, String value) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 12),
      child: Row(
        children: [
          Icon(icon, size: 20, color: Colors.grey[600]),
          const HSpace.medium(),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  label,
                  style: const TextStyle(
                    fontSize: 12,
                    fontWeight: FontWeight.w500,
                    color: Colors.grey,
                  ),
                ),
                Text(
                  value,
                  style: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  bool _hasContactInfo(Map<String, dynamic> profile) {
    return profile['email'] != null ||
        profile['github'] != null ||
        profile['telegram'] != null;
  }
}
