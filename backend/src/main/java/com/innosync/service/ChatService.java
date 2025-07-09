package com.innosync.service;

import com.innosync.model.ChatRoom;
import com.innosync.model.Message;
import com.innosync.model.Team;
import com.innosync.model.User;
import com.innosync.repository.ChatRoomRepository;
import com.innosync.repository.MessageRepository;
import com.innosync.repository.TeamRepository;
import com.innosync.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.HashSet;
import java.util.List;
import java.util.Optional;
import java.util.Set;

@Service
@RequiredArgsConstructor
public class ChatService {
    private final ChatRoomRepository chatRoomRepository;
    private final MessageRepository messageRepository;
    private final UserRepository userRepository;
    private final TeamRepository teamRepository;

    public List<ChatRoom> getChatRoomsForUser(Long userId) {
        Optional<User> userOpt = userRepository.findById(userId);
        if (userOpt.isEmpty()) return List.of();
        User user = userOpt.get();
        return chatRoomRepository.findAll().stream()
                .filter(room -> room.getParticipants().contains(user))
                .toList();
    }

    public List<Message> getMessagesForChatRoom(Long chatRoomId) {
        return messageRepository.findByChatRoomIdOrderByTimestampAsc(chatRoomId);
    }

    @Transactional
    public Message sendMessage(Long chatRoomId, Long senderId, String content) {
        ChatRoom chatRoom = chatRoomRepository.findById(chatRoomId)
                .orElseThrow(() -> new IllegalArgumentException("Chat room not found"));
        User sender = userRepository.findById(senderId)
                .orElseThrow(() -> new IllegalArgumentException("Sender not found"));
        Message message = Message.builder()
                .chatRoom(chatRoom)
                .sender(sender)
                .content(content)
                .timestamp(LocalDateTime.now())
                .build();
        return messageRepository.save(message);
    }

    @Transactional
    public ChatRoom getOrCreatePrivateChat(Long userId1, Long userId2) {
        User user1 = userRepository.findById(userId1)
                .orElseThrow(() -> new IllegalArgumentException("User not found"));
        User user2 = userRepository.findById(userId2)
                .orElseThrow(() -> new IllegalArgumentException("User not found"));
        Set<User> participants = new HashSet<>();
        participants.add(user1);
        participants.add(user2);
        // Find existing private chat
        return chatRoomRepository.findAll().stream()
                .filter(room -> !room.isTeamChat() && room.getParticipants().containsAll(participants) && room.getParticipants().size() == 2)
                .findFirst()
                .orElseGet(() -> {
                    ChatRoom newRoom = ChatRoom.builder()
                            .isTeamChat(false)
                            .participants(participants)
                            .build();
                    return chatRoomRepository.save(newRoom);
                });
    }

    @Transactional
    public ChatRoom getOrCreateTeamChat(Long teamId) {
        Team team = teamRepository.findById(teamId)
                .orElseThrow(() -> new IllegalArgumentException("Team not found"));
        // Find existing team chat
        return chatRoomRepository.findAll().stream()
                .filter(room -> room.isTeamChat() && room.getTeam() != null && room.getTeam().getId().equals(teamId))
                .findFirst()
                .orElseGet(() -> {
                    ChatRoom newRoom = ChatRoom.builder()
                            .isTeamChat(true)
                            .team(team)
                            .name(team.getName() + " Chat")
                            .participants(new HashSet<>(team.getMembers()))
                            .build();
                    return chatRoomRepository.save(newRoom);
                });
    }
} 