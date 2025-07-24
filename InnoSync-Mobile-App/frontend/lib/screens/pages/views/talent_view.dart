import 'package:flutter/material.dart';
import '../../../models/talent_model.dart';
import '../../../widgets/common/widgets.dart';

class TalentView extends StatelessWidget {
  final Talent talent;
  final VoidCallback? onBack;

  const TalentView({super.key, required this.talent, this.onBack});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.surface,
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            // Profile Image
            CircleAvatar(
              radius: 80,
              backgroundColor: Theme.of(context).colorScheme.surface,
              child: talent.profileImageUrl != null
                  ? ClipRRect(
                      borderRadius: BorderRadius.circular(80),
                      child: Image.network(
                        talent.profileImageUrl!,
                        fit: BoxFit.cover,
                        width: 160,
                        height: 160,
                      ),
                    )
                  : Icon(
                      Icons.person,
                      size: 80,
                      color: Theme.of(context).colorScheme.onSurface,
                    ),
            ),
            const VSpace(24),
            // Name
            Text(
              talent.name,
              style: Theme.of(
                context,
              ).textTheme.headlineSmall?.copyWith(fontSize: 28),
              textAlign: TextAlign.center,
            ),
            const VSpace(16),
            // Years and graduation type
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 12,
                    vertical: 6,
                  ),
                  decoration: BoxDecoration(
                    color: Theme.of(
                      context,
                    ).colorScheme.primary.withValues(alpha: 0.1),
                    borderRadius: BorderRadius.circular(16),
                  ),
                  child: Text(
                    '${talent.yearsOfExperience} years experience',
                    style: Theme.of(context).textTheme.bodyLarge?.copyWith(
                      color: Theme.of(context).colorScheme.primary,
                    ),
                  ),
                ),
                const HSpace.medium(),
                Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 12,
                    vertical: 6,
                  ),
                  decoration: BoxDecoration(
                    color: Theme.of(
                      context,
                    ).colorScheme.primary.withValues(alpha: 0.1),
                    borderRadius: BorderRadius.circular(16),
                  ),
                  child: Text(
                    talent.graduationType,
                    style: Theme.of(context).textTheme.bodyLarge?.copyWith(
                      color: Theme.of(context).colorScheme.primary,
                    ),
                  ),
                ),
              ],
            ),
            const VSpace(20),
            // Description
            Text(
              talent.description,
              style: Theme.of(
                context,
              ).textTheme.bodyLarge?.copyWith(fontSize: 16),
              textAlign: TextAlign.center,
            ),
            const VSpace(28),
            // Skills
            if (talent.skills.isNotEmpty) ...[
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Skills',
                  style: Theme.of(
                    context,
                  ).textTheme.titleMedium?.copyWith(fontSize: 16),
                ),
              ),
              const VSpace(8),
              Wrap(
                spacing: 10,
                runSpacing: 8,
                children: talent.skills
                    .map(
                      (skill) => Chip(
                        label: Text(
                          skill,
                          style: Theme.of(context).textTheme.bodyLarge
                              ?.copyWith(
                                color: Theme.of(context).colorScheme.onPrimary,
                              ),
                        ),
                        backgroundColor: Theme.of(context).colorScheme.primary,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(20),
                        ),
                      ),
                    )
                    .toList(),
              ),
              const VSpace(24),
            ],
            if (talent.positions.isNotEmpty) ...[
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Positions',
                  style: Theme.of(
                    context,
                  ).textTheme.titleMedium?.copyWith(fontSize: 16),
                ),
              ),
              const VSpace(8),
              Wrap(
                spacing: 10,
                runSpacing: 8,
                children: talent.positions
                    .map(
                      (position) => Chip(
                        label: Text(
                          position,
                          style: Theme.of(context).textTheme.bodyLarge
                              ?.copyWith(
                                color: Theme.of(context).colorScheme.onPrimary,
                              ),
                        ),
                        backgroundColor: Theme.of(context).colorScheme.primary,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(20),
                        ),
                      ),
                    )
                    .toList(),
              ),
              const VSpace(24),
            ],
            // Contact Information
            if (talent.email != null || talent.location != null) ...[
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Contact Information',
                  style: Theme.of(
                    context,
                  ).textTheme.titleMedium?.copyWith(fontSize: 16),
                ),
              ),
              const VSpace(8),
              if (talent.email != null) ...[
                ListTile(
                  leading: Icon(
                    Icons.email,
                    color: Theme.of(context).colorScheme.secondary,
                  ),
                  title: Text(talent.email!),
                  contentPadding: EdgeInsets.zero,
                ),
              ],
              if (talent.location != null) ...[
                ListTile(
                  leading: Icon(
                    Icons.location_on,
                    color: Theme.of(context).colorScheme.secondary,
                  ),
                  title: Text(talent.location!),
                  contentPadding: EdgeInsets.zero,
                ),
              ],
              const VSpace(32),
            ],
            // Contact button
            SizedBox(
              width: double.infinity,
              child: SubmitButton(
                text: 'Contact',
                onPressed: () async {
                  // TODO: Replace with actual API call if implemented
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('Contact request sent! (Demo)'),
                    ),
                  );
                },
                isLoading: false,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
