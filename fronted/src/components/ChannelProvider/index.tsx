import React from 'react'
import useGlobalRTM from 'hooks/useGlobalRTM'

const ChannelProvider: React.FC = () => {
  window.globalRTMApi = useGlobalRTM()

  return <></>
}

export default ChannelProvider
