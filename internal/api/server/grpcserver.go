package server

import (
	"context"
	"errors"
	"log"
	"net"
	"gocloudcamp/internal/api/proto"
	"gocloudcamp/internal/database"
	"gocloudcamp/internal/playlist"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	proto.UnimplementedApiServer
	playlist *playlist.Playlist
	database *database.Database
}

// Create new gRPC server
func New(pl *playlist.Playlist) *GRPCServer {
	return &GRPCServer{
		playlist: pl,
		database: database.NewDB(pl),
	}
}

// Starting server
func (s *GRPCServer) Start(newS *GRPCServer) {
	err := s.database.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterApiServer(server, newS)
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// Starting playing playlist
func (s *GRPCServer) Play(context.Context, *proto.Request) (*proto.Message, error) {
	s.database.GetAllSongs()
	s.playlist.Play()
	return &proto.Message{Message: "playlist starting"}, nil
}

// Pause playing playlist
func (s *GRPCServer) Pause(context.Context, *proto.Request) (*proto.Message, error) {
	s.playlist.Pause()
	return &proto.Message{Message: "playlist pause"}, nil
}

// Add new song to the end of playlist
func (s *GRPCServer) AddSong(ctx context.Context, ds *proto.DataSongRequest) (*proto.AddSongResponse, error) {
	name := ds.GetName()
	duration := ds.GetDuration()
	s.playlist.AddSong(name, int(duration))
	response := "song " + name + " added"
	return &proto.AddSongResponse{AddedSong: response}, nil
}

// Switching to the next song
func (s *GRPCServer) Next(context.Context, *proto.Request) (*proto.Message, error) {
	s.playlist.Next()
	return &proto.Message{Message: "Next song"}, nil
}

// Switching to the previous song
func (s *GRPCServer) Prev(context.Context, *proto.Request) (*proto.Message, error) {
	s.playlist.Prev()
	return &proto.Message{Message: "previous song"}, nil
}

// Stopping playing playlist
func (s *GRPCServer) Stop(context.Context, *proto.Request) (*proto.Message, error) {
	ok := s.playlist.IsEmpty()
	if !ok {
		s.database.SetPlaylist()
	}
	s.playlist.Stop()
	return &proto.Message{Message: "playlist stopped"}, nil
}

// Show now playing song if playlist is playind or show song on database searching by name
func (s *GRPCServer) ShowSong(ctx context.Context, sn *proto.SongName) (*proto.ShowSongResponse, error) {
	name := sn.GetName()
	ok := s.playlist.IsPlaying()
	if !ok {
		if name == "" {
			return nil, errors.New("Введите имя песни")
		}
		song, ok := s.database.GetSongByName(name)
		if !ok {
			return nil, errors.New("Песня не найдена")
		}
		return &proto.ShowSongResponse{Name: song.Name, Duration: int32(song.Duration)}, nil
	}

	songN, songD := s.playlist.ShowSong()
	return &proto.ShowSongResponse{Name: songN, Duration: int32(songD)}, nil
}

// Updating song on database if playlist is stopped
func (s *GRPCServer) UpdateSong(ctx context.Context, ds *proto.UpdateDataSongRequest) (*proto.Message, error) {
	oldName := ds.GetOldName()
	newName := ds.GetNewName()
	duration := ds.GetDuration()
	ok := s.playlist.IsPlaying()
	if ok {
		return nil, errors.New("Для редактирования песни проигрывание плейлиста должно быть остановлено")
	}
	err := s.database.UpdateSong(oldName, newName, int(duration))
	if err != nil {
		return nil, err
	}
	return &proto.Message{Message: "song is updated"}, nil
}

// Deleting song on database if playlist is stopped
func (s *GRPCServer) DeleteSong(ctx context.Context, sn *proto.SongName) (*proto.Message, error) {
	name := sn.GetName()
	ok := s.playlist.IsPlaying()
	if ok {
		return nil, errors.New("Для удаления песни проигрывание плейлиста должно быть остановлено")
	}
	err := s.database.DeleteSong(name)
	if err != nil {
		return nil, err
	}
	return &proto.Message{Message: "song is deleted"}, nil
}
