import 'package:flutter/material.dart';
import '../../../models/project_model.dart';
import '../../../widgets/common/widgets.dart';
import '../../../services/api_service.dart';
import '../../../utils/token_storage.dart';

class ProjectView extends StatelessWidget {
  final Project project;
  final VoidCallback? onBack;
  const ProjectView({super.key, required this.project, this.onBack});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.surface,
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            // Big square logo
            Container(
              width: 120,
              height: 120,
              decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.surface,
                borderRadius: BorderRadius.circular(24),
              ),
              child: project.logoUrl != null
                  ? ClipRRect(
                      borderRadius: BorderRadius.circular(24),
                      child: Image.network(project.logoUrl!, fit: BoxFit.cover),
                    )
                  : Icon(
                      Icons.apps,
                      size: 64,
                      color: Theme.of(context).colorScheme.onSurface,
                    ),
            ),
            const VSpace(24),
            // Bold Title
            Text(
              project.title,
              style: Theme.of(
                context,
              ).textTheme.headlineSmall?.copyWith(fontSize: 28),
              textAlign: TextAlign.center,
            ),
            const VSpace(16),
            // Description
            Text(
              project.description,
              style: Theme.of(
                context,
              ).textTheme.bodyLarge?.copyWith(fontSize: 16),
              textAlign: TextAlign.center,
            ),
            const VSpace(28),
            // Skills Needed
            if (project.skills.isNotEmpty) ...[
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Skills Needed',
                  style: Theme.of(
                    context,
                  ).textTheme.titleMedium?.copyWith(fontSize: 16),
                ),
              ),
              const VSpace(8),
              Wrap(
                spacing: 10,
                runSpacing: 8,
                children: project.skills
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
            // Available Positions
            if (project.positions.isNotEmpty) ...[
              Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Available Positions',
                  style: Theme.of(
                    context,
                  ).textTheme.titleMedium?.copyWith(fontSize: 16),
                ),
              ),
              const VSpace(8),
              Wrap(
                spacing: 10,
                runSpacing: 8,
                children: project.positions
                    .map(
                      (pos) => Chip(
                        label: Text(
                          pos,
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
              const VSpace(32),
            ],
            // Apply button
            SizedBox(
              width: double.infinity,
              child: SubmitButton(
                text: 'Apply',
                onPressed: () async {
                  final token = await getToken();
                  if (token == null) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text('Not authenticated!')),
                    );
                    return;
                  }
                  final api = ApiService();
                  final success = await api.applyToProject(
                    token: token,
                    projectId: project.id,
                  );
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text(
                        success ? 'Application sent!' : 'Failed to apply.',
                      ),
                    ),
                  );
                },
                isLoading: false,
              ),
            ),
            const VSpace(16),
            SizedBox(
              width: double.infinity,
              child: ElevatedButton.icon(
                icon: const Icon(Icons.person_add),
                label: const Text('Invite User'),
                onPressed: () {
                  showDialog(
                    context: context,
                    builder: (context) => _InviteUserDialog(projectId: project.id),
                  );
                },
              ),
            ),
            const VSpace(16),
            SizedBox(
              width: double.infinity,
              child: ElevatedButton.icon(
                icon: const Icon(Icons.add),
                label: const Text('Add Role'),
                onPressed: () {
                  showDialog(
                    context: context,
                    builder: (context) => _AddRoleDialog(projectId: project.id),
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}

// Диалог для приглашения пользователя по userId и projectRoleId
class _InviteUserDialog extends StatefulWidget {
  final int projectId;
  const _InviteUserDialog({required this.projectId});

  @override
  State<_InviteUserDialog> createState() => _InviteUserDialogState();
}

class _InviteUserDialogState extends State<_InviteUserDialog> {
  final _formKey = GlobalKey<FormState>();
  final _userIdController = TextEditingController();
  final _roleIdController = TextEditingController();
  final _messageController = TextEditingController();
  bool _isLoading = false;
  String? _error;
  String? _success;

  @override
  void dispose() {
    _userIdController.dispose();
    _roleIdController.dispose();
    _messageController.dispose();
    super.dispose();
  }

  Future<void> _sendInvitation() async {
    if (!_formKey.currentState!.validate()) return;
    setState(() {
      _isLoading = true;
      _error = null;
      _success = null;
    });
    try {
      final token = await getToken();
      if (token == null) throw Exception('Not authenticated');
      final api = ApiService();
      final ok = await api.createInvitation(
        token: token,
        projectRoleId: int.parse(_roleIdController.text),
        recipientId: int.parse(_userIdController.text),
        message: _messageController.text.isNotEmpty ? _messageController.text : null,
      );
      if (ok) {
        setState(() => _success = 'Invitation sent!');
        await Future.delayed(const Duration(seconds: 1));
        if (mounted) Navigator.of(context).pop();
      } else {
        setState(() => _error = 'Failed to send invitation');
      }
    } catch (e) {
      setState(() => _error = e.toString());
    } finally {
      setState(() => _isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Invite User to Project'),
      content: Form(
        key: _formKey,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextFormField(
              controller: _userIdController,
              decoration: const InputDecoration(labelText: 'User ID'),
              keyboardType: TextInputType.number,
              validator: (v) => v == null || v.isEmpty ? 'Enter user ID' : null,
            ),
            const SizedBox(height: 12),
            TextFormField(
              controller: _roleIdController,
              decoration: const InputDecoration(labelText: 'Project Role ID'),
              keyboardType: TextInputType.number,
              validator: (v) => v == null || v.isEmpty ? 'Enter role ID' : null,
            ),
            const SizedBox(height: 12),
            TextFormField(
              controller: _messageController,
              decoration: const InputDecoration(labelText: 'Message (optional)'),
              minLines: 1,
              maxLines: 3,
            ),
            if (_error != null)
              Padding(
                padding: const EdgeInsets.only(top: 8),
                child: Text(_error!, style: const TextStyle(color: Colors.red)),
              ),
            if (_success != null)
              Padding(
                padding: const EdgeInsets.only(top: 8),
                child: Text(_success!, style: const TextStyle(color: Colors.green)),
              ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: _isLoading ? null : () => Navigator.of(context).pop(),
          child: const Text('Cancel'),
        ),
        ElevatedButton(
          onPressed: _isLoading ? null : _sendInvitation,
          child: _isLoading
              ? const SizedBox(width: 16, height: 16, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
              : const Text('Send Invitation'),
        ),
      ],
    );
  }
}

// Диалог для создания project-role
class _AddRoleDialog extends StatefulWidget {
  final int projectId;
  const _AddRoleDialog({required this.projectId});

  @override
  State<_AddRoleDialog> createState() => _AddRoleDialogState();
}

class _AddRoleDialogState extends State<_AddRoleDialog> {
  final _formKey = GlobalKey<FormState>();
  final _roleNameController = TextEditingController();
  final _descriptionController = TextEditingController();
  final _skillsController = TextEditingController();
  final _countController = TextEditingController();
  bool _isLoading = false;
  String? _error;
  String? _success;

  @override
  void dispose() {
    _roleNameController.dispose();
    _descriptionController.dispose();
    _skillsController.dispose();
    _countController.dispose();
    super.dispose();
  }

  Future<void> _createRole() async {
    if (!_formKey.currentState!.validate()) return;
    setState(() {
      _isLoading = true;
      _error = null;
      _success = null;
    });
    try {
      final token = await getToken();
      if (token == null) throw Exception('Not authenticated');
      final api = ApiService();
      final ok = await api.createProjectRole(
        token: token,
        projectId: widget.projectId,
        roleName: _roleNameController.text,
        description: _descriptionController.text,
        skills: _skillsController.text.isNotEmpty
            ? _skillsController.text.split(',').map((s) => s.trim()).where((s) => s.isNotEmpty).toList()
            : null,
        count: _countController.text.isNotEmpty ? int.tryParse(_countController.text) : null,
      );
      if (ok) {
        setState(() => _success = 'Role created!');
        await Future.delayed(const Duration(seconds: 1));
        if (mounted) Navigator.of(context).pop();
        if (context.mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Role created!')),
          );
        }
      } else {
        setState(() => _error = 'Failed to create role');
      }
    } catch (e) {
      setState(() => _error = e.toString());
    } finally {
      setState(() => _isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Add Role to Project'),
      content: Form(
        key: _formKey,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextFormField(
              controller: _roleNameController,
              decoration: const InputDecoration(labelText: 'Role Name'),
              validator: (v) => v == null || v.isEmpty ? 'Enter role name' : null,
            ),
            const SizedBox(height: 12),
            TextFormField(
              controller: _descriptionController,
              decoration: const InputDecoration(labelText: 'Description (optional)'),
              minLines: 1,
              maxLines: 3,
            ),
            const SizedBox(height: 12),
            TextFormField(
              controller: _skillsController,
              decoration: const InputDecoration(labelText: 'Skills (comma separated, optional)'),
            ),
            const SizedBox(height: 12),
            TextFormField(
              controller: _countController,
              decoration: const InputDecoration(labelText: 'Count (optional)'),
              keyboardType: TextInputType.number,
            ),
            if (_error != null)
              Padding(
                padding: const EdgeInsets.only(top: 8),
                child: Text(_error!, style: const TextStyle(color: Colors.red)),
              ),
            if (_success != null)
              Padding(
                padding: const EdgeInsets.only(top: 8),
                child: Text(_success!, style: const TextStyle(color: Colors.green)),
              ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: _isLoading ? null : () => Navigator.of(context).pop(),
          child: const Text('Cancel'),
        ),
        ElevatedButton(
          onPressed: _isLoading ? null : _createRole,
          child: _isLoading
              ? const SizedBox(width: 16, height: 16, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
              : const Text('Create Role'),
        ),
      ],
    );
  }
}
