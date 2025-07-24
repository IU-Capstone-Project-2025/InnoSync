
class Invitation {
  final int id;
  final String? message;
  final int projectId;
  final int projectRoleId;
  final String projectTitle;
  final int recipientId;
  final String recipientName;
  final DateTime? respondedAt;
  final String roleName;
  final DateTime sentAt;
  final String invitationStatus;

  Invitation({
    required this.id,
    this.message,
    required this.projectId,
    required this.projectRoleId,
    required this.projectTitle,
    required this.recipientId,
    required this.recipientName,
    this.respondedAt,
    required this.roleName,
    required this.sentAt,
    required this.invitationStatus,
  });

  factory Invitation.fromJson(Map<String, dynamic> json) {
    return Invitation(
      id: json['id'],
      message: json['message'] as String?,
      projectId: json['project_id'],
      projectRoleId: json['project_role_id'],
      projectTitle: json['project_title'],
      recipientId: json['recipient_id'],
      recipientName: json['recipient_name'],
      respondedAt: json['responded_at'] != null
          ? DateTime.parse(json['responded_at'])
          : null,
      roleName: json['role_name'],
      sentAt: DateTime.parse(json['sent_at']),
      invitationStatus: json['invitation_status'],
    );
  }
}