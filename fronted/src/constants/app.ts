const App = {
  Host: 'https://italking.tomotoes.com',
  MaxMemberShowLength: 3,
  FetchRoomInterval: 1e3 * 60,
  AgoraAppId: process.env.REACT_APP_AGORA_ID!,
  SentryDSN: process.env.REACT_APP_SENTRY_DSN!,
  ChatIdPrefix: 'cid-',
  MaxMessageLength: 500,
  GlobalChannelName: 'ITalking-Channel',
  DescriptionMaxLength: 30,
  NameMaxLength: 10,
  NameMinLength: 2,
  PasswordMaxLength: 15,
  PasswordMinLength: 6,
  RoomNameMaxLength: 15,
  RoomNameMinLength: 2,
  RoomDescriptionMaxLength: 80,
  RequestSpeakingDisplayTime: 3e3,
  SoundEffectVolume: 0.7,
  IsProduction: process.env.NODE_ENV === 'production',
  IssuesUrl: 'https://github.com/Tomotoes/ITalking/issues',
  AuthorUrl: 'https://tomotoes.com'
}

export default App
