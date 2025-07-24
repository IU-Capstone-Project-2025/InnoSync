import 'package:flutter/material.dart';
import 'package:frontend/theme/colors.dart';
import 'package:frontend/theme/dimensions.dart';
import 'package:frontend/theme/text_styles.dart';
import 'package:frontend/widgets/common/widgets.dart';
import 'package:frontend/widgets/login/widgets.dart';
import 'dart:io';
import '../../../services/api_service.dart';
import '../../../utils/token_storage.dart';

class ProjectCreate extends StatefulWidget {
  final VoidCallback? onProjectCreated;
  const ProjectCreate({super.key, this.onProjectCreated});

  @override
  State<ProjectCreate> createState() => _ProjectCreateState();
}

class _ProjectCreateState extends State<ProjectCreate> {
  File? _logoFile;
  final TextEditingController _titleController = TextEditingController();
  final TextEditingController _descriptionController = TextEditingController();
  final List<String> _skills = [];
  final List<String> _positions = [];
  final TextEditingController _skillsController = TextEditingController();
  final TextEditingController _teamSizeController = TextEditingController();
  final TextEditingController _positionsController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: MainAppBar(title: 'Create Project'),
      body: Stack(
        children: [
          Padding(
            padding: const EdgeInsets.only(
              bottom: 70,
            ), // leave space for floating chips
            child: Column(
              children: [
                const VSpace.md(),
                ProfileImagePicker(
                  profileImage: _logoFile,
                  onImagePicked: (file) => setState(() => _logoFile = file),
                  backgroundColor: AppColors.profilePickerBackground,
                ),
                const VSpace.lg(),
                TextField(
                  controller: _titleController,
                  decoration: InputDecoration(
                    contentPadding: const EdgeInsets.symmetric(
                      horizontal: 16,
                      vertical: 12,
                    ),
                    labelText: 'Title',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(
                        AppDimensions.radiusLg,
                      ),
                    ),
                  ),
                ),
                const VSpace.md(),
                TextField(
                  controller: _descriptionController,
                  decoration: InputDecoration(
                    contentPadding: const EdgeInsets.symmetric(
                      horizontal: 16,
                      vertical: 12,
                    ),
                    labelText: 'Description',
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(
                        AppDimensions.radiusLg,
                      ),
                    ),
                  ),
                ),
                const VSpace.md(),
                Row(
                  children: [
                    Expanded(
                      child: TextField(
                        controller: _skillsController,
                        decoration: InputDecoration(
                          contentPadding: const EdgeInsets.symmetric(
                            horizontal: 16,
                            vertical: 12,
                          ),
                          labelText: 'Skills',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(
                              AppDimensions.radiusLg,
                            ),
                          ),
                          suffixIcon: IconButton(
                            icon: const Icon(Icons.add),
                            onPressed: () {
                              final value = _skillsController.text.trim();
                              if (value.isNotEmpty) {
                                setState(() {
                                  _skills.add(value);
                                  _skillsController.clear();
                                });
                              }
                            },
                            iconSize: 24,
                          ),
                        ),
                        onSubmitted: (value) {
                          if (value.trim().isNotEmpty) {
                            setState(() {
                              _skills.add(value.trim());
                              _skillsController.clear();
                            });
                          }
                        },
                      ),
                    ),
                    const VSpace.small(),
                  ],
                ),
                const VSpace.md(),
                TextField(
                  controller: _teamSizeController,
                  decoration: InputDecoration(labelText: 'Team Size'),
                ),
                if (_skills.isNotEmpty)
                  Positioned(
                    left: 0,
                    right: 0,
                    bottom: 0,
                    child: Container(
                      color: Colors.transparent,
                      padding: const EdgeInsets.symmetric(
                        horizontal: 16,
                        vertical: 8,
                      ),
                      child: SingleChildScrollView(
                        scrollDirection: Axis.horizontal,
                        child: Row(
                          children: _skills
                              .map(
                                (keyword) => Padding(
                                  padding: const EdgeInsets.only(right: 8.0),
                                  child: Chip(
                                    label: Text(
                                      keyword,
                                      style: AppTextStyles.bodyLarge.copyWith(
                                        color: AppColors.textOnPrimary,
                                      ),
                                    ),
                                    backgroundColor: AppColors.primary,
                                    shape: RoundedRectangleBorder(
                                      borderRadius: BorderRadius.circular(24),
                                    ),
                                    deleteIcon: Container(
                                      decoration: const BoxDecoration(
                                        color: AppColors.textOnPrimary,
                                        shape: BoxShape.circle,
                                      ),
                                      child: const Icon(
                                        Icons.close,
                                        size: 18,
                                        color: AppColors.primary,
                                      ),
                                    ),
                                    onDeleted: () =>
                                        setState(() => _skills.remove(keyword)),
                                  ),
                                ),
                              )
                              .toList(),
                        ),
                      ),
                    ),
                  ),
                const VSpace.md(),
                Row(
                  children: [
                    Expanded(
                      child: TextField(
                        controller: _positionsController,
                        decoration: InputDecoration(
                          contentPadding: const EdgeInsets.symmetric(
                            horizontal: 16,
                            vertical: 12,
                          ),
                          labelText: 'Positions',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(
                              AppDimensions.radiusLg,
                            ),
                          ),
                          suffixIcon: IconButton(
                            icon: const Icon(Icons.add),
                            onPressed: () {
                              final value = _positionsController.text.trim();
                              if (value.isNotEmpty) {
                                setState(() {
                                  _positions.add(value);
                                  _positionsController.clear();
                                });
                              }
                            },
                            iconSize: 24,
                          ),
                        ),
                        onSubmitted: (value) {
                          if (value.trim().isNotEmpty) {
                            setState(() {
                              _positions.add(value.trim());
                              _positionsController.clear();
                            });
                          }
                        },
                      ),
                    ),
                    const VSpace.small(),
                  ],
                ),
                if (_skills.isNotEmpty)
                  Positioned(
                    left: 0,
                    right: 0,
                    bottom: 0,
                    child: Container(
                      color: Colors.transparent,
                      padding: const EdgeInsets.symmetric(
                        horizontal: 16,
                        vertical: 8,
                      ),
                      child: SingleChildScrollView(
                        scrollDirection: Axis.horizontal,
                        child: Row(
                          children: _positions
                              .map(
                                (keyword) => Padding(
                                  padding: const EdgeInsets.only(right: 8.0),
                                  child: Chip(
                                    label: Text(
                                      keyword,
                                      style: AppTextStyles.bodyLarge.copyWith(
                                        color: AppColors.textOnPrimary,
                                      ),
                                    ),
                                    backgroundColor: AppColors.primary,
                                    shape: RoundedRectangleBorder(
                                      borderRadius: BorderRadius.circular(24),
                                    ),
                                    deleteIcon: Container(
                                      decoration: const BoxDecoration(
                                        color: AppColors.textOnPrimary,
                                        shape: BoxShape.circle,
                                      ),
                                      child: const Icon(
                                        Icons.close,
                                        size: 18,
                                        color: AppColors.primary,
                                      ),
                                    ),
                                    onDeleted: () => setState(
                                      () => _positions.remove(keyword),
                                    ),
                                  ),
                                ),
                              )
                              .toList(),
                        ),
                      ),
                    ),
                  ),
                const Spacer(),
                ElevatedButton(
                  onPressed: () async {
                    final token = await getToken();
                    if (token == null) {
                      if (!mounted) return;
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(content: Text('Not authenticated!')),
                      );
                      return;
                    }
                    final success = await ApiService.createProject(
                      token,
                      _titleController.text,
                      _descriptionController.text,
                      int.tryParse(_teamSizeController.text),
                    );
                    if (!mounted) return;
                    if (success) {
                      widget.onProjectCreated?.call();
                      if (!mounted) return;
                      Navigator.pop(context);
                    } else {
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(
                          content: Text('Failed to create project.'),
                        ),
                      );
                    }
                  },
                  child: const Text('Create Project'),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
