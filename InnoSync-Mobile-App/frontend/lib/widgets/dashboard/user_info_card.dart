import 'package:flutter/material.dart';
import '../common/widgets.dart';

class UserInfoCard extends StatelessWidget {
  final String? name;
  final String? email;
  final String? bio;
  final String? positions;
  final String? technologies;
  final String? avatarUrl;
  final double width;
  final double avatarRadius;
  final VoidCallback? onTap;

  const UserInfoCard({
    super.key,
    this.name,
    this.email,
    this.bio,
    this.positions,
    this.technologies,
    this.avatarUrl,
    this.width = 250,
    this.avatarRadius = 50,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return GestureDetector(
      onTap: onTap,
      child: Card(
        elevation: 2,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
        child: Container(
          decoration: BoxDecoration(
            color: theme.colorScheme.surface,
            borderRadius: BorderRadius.circular(16),
            boxShadow: [
              BoxShadow(
                color: theme.colorScheme.shadow.withValues(alpha: 0.1),
                blurRadius: 8,
                offset: const Offset(0, 2),
              ),
            ],
          ),
          padding: const EdgeInsets.all(20),
          child: Row(
            children: [
              CircleAvatar(
                radius: 30,
                backgroundImage: avatarUrl != null
                    ? NetworkImage(avatarUrl!)
                    : null,
                child: avatarUrl == null
                    ? Text(
                        name?.isNotEmpty ?? false
                            ? name![0].toUpperCase()
                            : '?',
                        style: theme.textTheme.titleLarge?.copyWith(
                          color: theme.colorScheme.onPrimary,
                          fontWeight: FontWeight.bold,
                        ),
                      )
                    : null,
              ),
              const HSpace.medium(),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      name ?? '',
                      style: theme.textTheme.titleMedium?.copyWith(
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    const VSpace.small(),
                    Text(
                      email ?? '',
                      style: theme.textTheme.bodyMedium?.copyWith(
                        color: theme.colorScheme.onSurface.withValues(
                          alpha: 0.7,
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
    );
  }
}
