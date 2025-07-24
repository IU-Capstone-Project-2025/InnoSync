import 'package:flutter/material.dart';
import 'package:frontend/models/invitation_model.dart';
import 'package:frontend/widgets/common/spacing.dart';
import 'package:frontend/widgets/common/widgets.dart';

class InvitationView extends StatelessWidget {
  final Invitation invitation;
  const InvitationView({super.key, required this.invitation});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    // Get initials for avatar
    String initials = invitation.recipientName.isNotEmpty
        ? invitation.recipientName
              .trim()
              .split(' ')
              .map((e) => e.isNotEmpty ? e[0] : '')
              .take(2)
              .join()
              .toUpperCase()
        : '?';
    return Scaffold(
      backgroundColor: Theme.of(context).scaffoldBackgroundColor,
      body: Center(
        child: SingleChildScrollView(
          padding: const EdgeInsets.all(24),
          child: Card(
            elevation: 4,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(24),
            ),
            margin: const EdgeInsets.symmetric(vertical: 16, horizontal: 0),
            child: Padding(
              padding: const EdgeInsets.all(28),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  // Avatar
                  CircleAvatar(
                    radius: 36,
                    backgroundColor: theme.colorScheme.primary.withValues(
                      alpha: 0.15,
                    ),
                    child: Text(
                      initials,
                      style: theme.textTheme.titleLarge?.copyWith(
                        color: theme.colorScheme.primary,
                        fontWeight: FontWeight.bold,
                        fontSize: 32,
                      ),
                    ),
                  ),
                  const VSpace(24),
                  Text(
                    invitation.recipientName,
                    style: theme.textTheme.titleLarge?.copyWith(fontSize: 22),
                    textAlign: TextAlign.center,
                  ),
                  const VSpace(16),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Container(
                        padding: const EdgeInsets.symmetric(
                          horizontal: 12,
                          vertical: 6,
                        ),
                        decoration: BoxDecoration(
                          color: theme.colorScheme.primary.withValues(
                            alpha: 0.12,
                          ),
                          borderRadius: BorderRadius.circular(14),
                        ),
                        child: Row(
                          children: [
                            Icon(
                              Icons.work_outline,
                              size: 18,
                              color: theme.colorScheme.primary,
                            ),
                            const SizedBox(width: 4),
                            Text(
                              invitation.roleName,
                              style: theme.textTheme.bodyLarge?.copyWith(
                                color: theme.colorScheme.primary,
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                          ],
                        ),
                      ),
                      const SizedBox(width: 12),
                      Container(
                        padding: const EdgeInsets.symmetric(
                          horizontal: 12,
                          vertical: 6,
                        ),
                        decoration: BoxDecoration(
                          color: theme.colorScheme.primary.withValues(
                            alpha: 0.12,
                          ),
                          borderRadius: BorderRadius.circular(14),
                        ),
                        child: Row(
                          children: [
                            Icon(
                              Icons.folder_open,
                              size: 18,
                              color: theme.colorScheme.primary,
                            ),
                            const SizedBox(width: 4),
                            Text(
                              invitation.projectTitle,
                              style: theme.textTheme.bodyLarge?.copyWith(
                                color: theme.colorScheme.primary,
                                fontSize: 14,
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                  const VSpace(24),
                  Text(
                    invitation.message ?? '',
                    style: theme.textTheme.bodyLarge?.copyWith(
                      color: theme.textTheme.bodyLarge?.color,
                      fontSize: 17,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  const VSpace(24),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(
                        Icons.calendar_today,
                        size: 18,
                        color: theme.colorScheme.onSurface.withValues(
                          alpha: 0.5,
                        ),
                      ),
                      const SizedBox(width: 6),
                      Text(
                        invitation.sentAt.toString(),
                        style: theme.textTheme.bodyLarge?.copyWith(
                          color: theme.colorScheme.onSurface.withValues(
                            alpha: 0.5,
                          ),
                          fontSize: 14,
                        ),
                      ),
                    ],
                  ),
                  const VSpace(16),
                  Container(
                    padding: const EdgeInsets.symmetric(
                      horizontal: 16,
                      vertical: 8,
                    ),
                    decoration: BoxDecoration(
                      color: invitation.invitationStatus == 'Accepted'
                          ? theme.colorScheme.primary.withValues(alpha: 0.12)
                          : invitation.invitationStatus == 'Declined'
                          ? theme.colorScheme.error.withValues(alpha: 0.12)
                          : theme.colorScheme.primary.withValues(alpha: 0.10),
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Text(
                      invitation.invitationStatus,
                      style: theme.textTheme.bodyLarge?.copyWith(
                        color: invitation.invitationStatus == 'Accepted'
                            ? theme.colorScheme.primary
                            : invitation.invitationStatus == 'Declined'
                            ? theme.colorScheme.error
                            : theme.colorScheme.primary,
                        fontWeight: FontWeight.w600,
                        fontSize: 15,
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
