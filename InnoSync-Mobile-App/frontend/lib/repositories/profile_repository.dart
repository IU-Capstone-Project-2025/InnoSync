import '../services/api_service.dart';
import 'dart:io';

class ProfileRepository {
  final ApiService apiService;

  ProfileRepository({required this.apiService});

  Future<Map<String, dynamic>> fetchProfile(String token) {
    return apiService.fetchProfile(token);
  }

  Future<bool> createProfile({
    required String token,
    required String name,
    required String email,
    String? telegram,
    String? github,
    String? bio,
    String? position,
    String? education,
    String? expertise,
    String? expertiseLevel,
    List<String>? technologies,
    List<Map<String, dynamic>>? workExperience,
    String? resumeUrl,
    File? profileImage,
  }) async {
    return await apiService.createProfile(
      token: token,
      name: name,
      email: email,
      telegram: telegram,
      github: github,
      bio: bio,
      position: position,
      education: education,
      expertise: expertise,
      expertiseLevel: expertiseLevel,
      technologies: technologies,
      workExperience: workExperience,
      resumeUrl: resumeUrl,
      profileImage: profileImage,
    );
  }
} 