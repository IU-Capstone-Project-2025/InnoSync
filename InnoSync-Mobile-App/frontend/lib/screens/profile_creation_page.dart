import 'package:flutter/material.dart';
import 'dart:io';
import '../widgets/common/widgets.dart';
import '../widgets/login/widgets.dart';
import '../theme/text_styles.dart';
import '../theme/dimensions.dart';
import '../theme/app_theme.dart';
import '../repositories/profile_repository.dart';
import '../utils/token_storage.dart';
import '../services/api_service.dart';
import '../utils/ui_helpers.dart';
import 'package:file_picker/file_picker.dart';
import 'dart:convert'; // Added for jsonDecode
import 'package:http/http.dart' as http; // Added for http

class ProfileCreationPage extends StatefulWidget {
  final Map<String, dynamic>? profileData;
  const ProfileCreationPage({super.key, this.profileData});

  @override
  State<ProfileCreationPage> createState() => _ProfileCreationPageState();
}

class _ProfileCreationPageState extends State<ProfileCreationPage> {
  final TextEditingController _fullNameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _telegramController = TextEditingController();
  final TextEditingController _githubController = TextEditingController();
  final TextEditingController _bioController = TextEditingController();
  final TextEditingController _positionController = TextEditingController();
  final TextEditingController _expertiseController = TextEditingController();
  File? _profileImage;
  File? _resumeFile;
  String? _resumeUrl;

  // Education & Expertise Level enums
  static const List<String> _educationOptions = [
    'NO_DEGREE',
    'BACHELOR',
    'MASTER',
    'PHD',
  ];
  String? _selectedEducation;

  static const List<String> _expertiseLevelOptions = [
    'ENTRY',
    'JUNIOR',
    'MID',
    'SENIOR',
    'RESEARCHER',
  ];
  String? _selectedExpertiseLevel;

  // Technologies (multi-select)
  final List<String> _allTechnologies = [
    'Flutter',
    'Dart',
    'Go',
    'React',
    'Node.js',
    'Python',
    'Java',
    'AWS',
    'Docker',
    'Kubernetes',
    'Figma',
    'Sketch',
    'Adobe XD',
    'Prototyping',
  ];
  final List<String> _selectedTechnologies = [];

  // Work Experience
  final List<Map<String, dynamic>> _workExperiences = [];

  @override
  void initState() {
    super.initState();
    if (widget.profileData != null) {
      final p = widget.profileData!;
      _fullNameController.text = p['name'] ?? '';
      _emailController.text = p['email'] ?? '';
      _telegramController.text = p['telegram'] ?? '';
      _githubController.text = p['github'] ?? '';
      _bioController.text = p['bio'] ?? '';
      _positionController.text = p['position'] ?? '';
      _selectedEducation = p['education'] ?? null;
      _expertiseController.text = p['expertise'] ?? '';
      _selectedExpertiseLevel = p['expertise_level'] ?? null;
      _selectedTechnologies.clear();
      if (p['technologies'] is List) {
        _selectedTechnologies.addAll(
          (p['technologies'] as List)
              .map((t) => t['name']?.toString() ?? '')
              .where((t) => t.isNotEmpty)
              .cast<String>(),
        );
      }
      _workExperiences.clear();
      if (p['work_experiences'] is List) {
        for (final we in p['work_experiences']) {
          _workExperiences.add({
            'position': we['position'] ?? '',
            'company': we['company'] ?? '',
            'startDate': (we['start_date'] ?? '').toString().substring(0, 10),
            'endDate': (we['end_date'] ?? '').toString().isNotEmpty
                ? we['end_date'].toString().substring(0, 10)
                : '',
            'description': we['description'] ?? '',
          });
        }
      }
      _resumeUrl = p['resume_url'];
    }
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    final args =
        ModalRoute.of(context)?.settings.arguments as Map<String, dynamic>?;
    if (args != null) {
      _emailController.text = args['email'] ?? '';
      _fullNameController.text = args['name'] ?? '';
    }
  }

  @override
  void dispose() {
    _fullNameController.dispose();
    _emailController.dispose();
    _telegramController.dispose();
    _githubController.dispose();
    _bioController.dispose();
    _positionController.dispose();
    _expertiseController.dispose();
    super.dispose();
  }

  void _onImagePicked(File image) {
    setState(() {
      _profileImage = image;
    });
  }

  Future<void> _pickResume() async {
    final result = await FilePicker.platform.pickFiles(type: FileType.any);
    if (result != null && result.files.single.path != null) {
      setState(() {
        _resumeFile = File(result.files.single.path!);
      });
    }
  }

  Future<void> _uploadResume() async {
    final token = await getToken();
    if (_resumeFile == null || token == null) return;
    final request = http.MultipartRequest(
      'POST',
      Uri.parse('${ApiService.baseUrl}/profile/resume'),
    );
    request.headers['Authorization'] = 'Bearer $token';
    request.files.add(
      await http.MultipartFile.fromPath('resume', _resumeFile!.path),
    );
    final response = await request.send();
    if (response.statusCode == 200) {
      final respStr = await response.stream.bytesToString();
      final data = jsonDecode(respStr);
      setState(() {
        _resumeUrl = data['url'] as String?;
      });
      UIHelpers.showSuccess(context, 'Resume uploaded!');
    } else {
      UIHelpers.showError(context, 'Resume upload failed');
    }
  }

  void _addWorkExperience() {
    setState(() {
      _workExperiences.add({
        'position': '',
        'company': '',
        'startDate': '',
        'endDate': '',
        'description': '',
      });
    });
  }

  void _removeWorkExperience(int index) {
    setState(() {
      _workExperiences.removeAt(index);
    });
  }

  List<Map<String, dynamic>> _prepareWorkExperiences() {
    final List<Map<String, dynamic>> result = [];
    for (final exp in _workExperiences) {
      if ((exp['position'] ?? '').trim().isEmpty ||
          (exp['company'] ?? '').trim().isEmpty ||
          (exp['startDate'] ?? '').trim().isEmpty) {
        continue; // пропускать невалидные
      }
      result.add({
        'position': exp['position'],
        'company': exp['company'],
        'start_date': exp['startDate'],
        'end_date': (exp['endDate'] ?? '').trim().isEmpty
            ? null
            : exp['endDate'],
        'description': (exp['description'] ?? '').trim().isEmpty
            ? null
            : exp['description'],
      });
    }
    return result;
  }

  @override
  Widget build(BuildContext context) {
    final apiService = ApiService();
    final profileRepository = ProfileRepository(apiService: apiService);
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.background,
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => Navigator.of(context).pop(),
        ),
        title: const Text('Edit Profile'),
        backgroundColor: Theme.of(context).colorScheme.background,
        elevation: 0,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(AppDimensions.paddingXl),
        child: Column(
          children: [
            ProfileImagePicker(
              profileImage: _profileImage,
              onImagePicked: _onImagePicked,
              backgroundColor: Theme.of(context).colorScheme.surface,
            ),
            const VSpace.lg(),
            _buildInputField('Full Name', _fullNameController, readOnly: true),
            const VSpace.lg(),
            _buildInputField('Email', _emailController, readOnly: true),
            const VSpace.lg(),
            _buildInputField('Telegram', _telegramController),
            const VSpace.lg(),
            _buildInputField('Github', _githubController),
            const VSpace.lg(),
            _buildInputField('Bio', _bioController, maxLines: 4),
            const VSpace.lg(),
            _buildInputField('Position', _positionController),
            const VSpace.lg(),
            // Education dropdown
            DropdownButtonFormField<String>(
              value: _selectedEducation,
              items: _educationOptions
                  .map((e) => DropdownMenuItem(value: e, child: Text(e)))
                  .toList(),
              onChanged: (val) => setState(() => _selectedEducation = val),
              decoration: InputDecoration(
                contentPadding: const EdgeInsets.symmetric(
                  horizontal: 16,
                  vertical: 12,
                ),
                labelText: 'Education',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
                ),
              ),
            ),
            const VSpace.lg(),
            _buildInputField('Expertise', _expertiseController),
            const VSpace.lg(),
            // Expertise Level dropdown
            DropdownButtonFormField<String>(
              value: _selectedExpertiseLevel,
              items: _expertiseLevelOptions
                  .map((e) => DropdownMenuItem(value: e, child: Text(e)))
                  .toList(),
              onChanged: (val) => setState(() => _selectedExpertiseLevel = val),
              decoration: InputDecoration(
                contentPadding: const EdgeInsets.symmetric(
                  horizontal: 16,
                  vertical: 12,
                ),
                labelText: 'Expertise Level',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
                ),
              ),
            ),
            const VSpace.lg(),
            // Technologies multi-select
            Align(
              alignment: Alignment.centerLeft,
              child: Text('Technologies', style: AppTextStyles.bodyLarge),
            ),
            Wrap(
              spacing: 8,
              runSpacing: 8,
              children: _allTechnologies
                  .map(
                    (tech) => FilterChip(
                      label: Text(tech),
                      selected: _selectedTechnologies.contains(tech),
                      onSelected: (selected) {
                        setState(() {
                          if (selected) {
                            _selectedTechnologies.add(tech);
                          } else {
                            _selectedTechnologies.remove(tech);
                          }
                        });
                      },
                    ),
                  )
                  .toList(),
            ),
            const VSpace.lg(),
            // Work Experience
            Align(
              alignment: Alignment.centerLeft,
              child: Text('Work Experience', style: AppTextStyles.bodyLarge),
            ),
            ..._workExperiences.asMap().entries.map((entry) {
              final i = entry.key;
              final exp = entry.value;
              return Card(
                margin: const EdgeInsets.symmetric(vertical: 8),
                child: Padding(
                  padding: const EdgeInsets.all(8),
                  child: Column(
                    children: [
                      _buildWorkExpField('Position', exp, 'position', i),
                      const VSpace.sm(),
                      _buildWorkExpField('Company', exp, 'company', i),
                      const VSpace.sm(),
                      _buildWorkExpField(
                        'Start Date (YYYY-MM-DD)',
                        exp,
                        'startDate',
                        i,
                      ),
                      const VSpace.sm(),
                      _buildWorkExpField(
                        'End Date (YYYY-MM-DD)',
                        exp,
                        'endDate',
                        i,
                      ),
                      const VSpace.sm(),
                      _buildWorkExpField(
                        'Description',
                        exp,
                        'description',
                        i,
                        maxLines: 2,
                      ),
                      Align(
                        alignment: Alignment.centerRight,
                        child: IconButton(
                          icon: const Icon(Icons.delete, color: Colors.red),
                          onPressed: () => _removeWorkExperience(i),
                        ),
                      ),
                    ],
                  ),
                ),
              );
            }),
            Align(
              alignment: Alignment.centerLeft,
              child: TextButton.icon(
                icon: const Icon(Icons.add),
                label: const Text('Add Work Experience'),
                onPressed: _addWorkExperience,
              ),
            ),
            const VSpace.lg(),
            // Resume upload
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ElevatedButton(
                  onPressed: () async {
                    await _pickResume();
                    await _uploadResume();
                  },
                  child: const Text('Upload Resume'),
                ),
                if (_resumeFile != null) ...[
                  const SizedBox(height: 8),
                  Text('Selected file: ${_resumeFile!.path.split('/').last}'),
                ],
                if (_resumeUrl != null) ...[
                  const SizedBox(height: 8),
                  Text('Resume uploaded!'),
                ],
              ],
            ),
            const VSpace.xxl(),
            ElevatedButton(
              style: AppTheme.primaryButtonStyle,
              onPressed: () async {
                final token = await getToken();
                final success = await profileRepository.createProfile(
                  token: token!,
                  name: _fullNameController.text,
                  email: _emailController.text,
                  telegram: _telegramController.text.isEmpty
                      ? null
                      : _telegramController.text,
                  github: _githubController.text.isEmpty
                      ? null
                      : _githubController.text,
                  bio: _bioController.text.isEmpty ? null : _bioController.text,
                  position: _positionController.text.isEmpty
                      ? null
                      : _positionController.text,
                  education: _selectedEducation,
                  expertise: _expertiseController.text.isEmpty
                      ? null
                      : _expertiseController.text,
                  expertiseLevel: _selectedExpertiseLevel,
                  technologies: _selectedTechnologies,
                  workExperience: _prepareWorkExperiences(),
                  resumeUrl: _resumeUrl,
                  profileImage: _profileImage,
                );
                if (success) {
                  await saveToken(token);
                  if (!mounted) return;
                  Navigator.pushReplacementNamed(context, '/main');
                } else {
                  if (!mounted) return;
                  UIHelpers.showError(context, 'Failed to create profile');
                }
              },
              child: const Text('Create Profile'),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildInputField(
    String label,
    TextEditingController controller, {
    bool readOnly = false,
    int maxLines = 1,
  }) {
    return TextField(
      controller: controller,
      readOnly: readOnly,
      maxLines: maxLines,
      decoration: InputDecoration(
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 12,
        ),
        labelText: label,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
        ),
      ),
    );
  }

  Widget _buildWorkExpField(
    String label,
    Map<String, dynamic> exp,
    String key,
    int index, {
    int maxLines = 1,
  }) {
    return TextField(
      controller: TextEditingController(text: exp[key] ?? '')
        ..selection = TextSelection.collapsed(offset: (exp[key] ?? '').length),
      maxLines: maxLines,
      decoration: InputDecoration(
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 12,
        ),
        labelText: label,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppDimensions.radiusLg),
        ),
      ),
      onChanged: (val) => setState(() => _workExperiences[index][key] = val),
    );
  }
}
