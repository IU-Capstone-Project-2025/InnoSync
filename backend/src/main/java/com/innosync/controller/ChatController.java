package com.innosync.controller;

import com.innosync.model.ChatRoom;
import com.innosync.model.Message;
import com.innosync.service.ChatService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/chats")
@RequiredArgsConstructor
public class ChatController {
    private final ChatService chatService;

    @GetMapping("/rooms/user/{userId}")
    public ResponseEntity<List<ChatRoom>> getChatRoomsForUser(@PathVariable Long userId) {
        return ResponseEntity.ok(chatService.getChatRoomsForUser(userId));
    }

    @GetMapping("/rooms/{chatRoomId}/messages")
    public ResponseEntity<List<Message>> getMessagesForChatRoom(@PathVariable Long chatRoomId) {
        return ResponseEntity.ok(chatService.getMessagesForChatRoom(chatRoomId));
    }

    @PostMapping("/rooms/{chatRoomId}/messages")
    public ResponseEntity<Message> sendMessage(@PathVariable Long chatRoomId, @RequestParam Long senderId, @RequestParam String content) {
        return ResponseEntity.ok(chatService.sendMessage(chatRoomId, senderId, content));
    }

    @PostMapping("/private")
    public ResponseEntity<ChatRoom> getOrCreatePrivateChat(@RequestParam Long userId1, @RequestParam Long userId2) {
        return ResponseEntity.ok(chatService.getOrCreatePrivateChat(userId1, userId2));
    }

    @PostMapping("/team")
    public ResponseEntity<ChatRoom> getOrCreateTeamChat(@RequestParam Long teamId) {
        return ResponseEntity.ok(chatService.getOrCreateTeamChat(teamId));
    }
} 