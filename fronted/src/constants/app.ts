const isProduction = process.env.NODE_ENV === 'production'
const App = {
  Host: 'https://italking.tomotoes.com',
  MaxMemberShowLength: 3,
  FetchRoomInterval: 1e3 * 60,
  AgoraAppId: 'df244aeb1e3447608f1c72025da73b85',
  SentryDSN: 'https://8a13ff80301a410ab007b1ed694672e0@o717267.ingest.sentry.io/5779797',
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
  isProduction
}

export default App
