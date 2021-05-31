import React from 'react'
import { FallbackProps } from 'components/ErrorBoundary'

const Fallback: React.FC<FallbackProps> = props => {
  const { error } = props
  return <div>Error: {error.message}</div>
}

export default Fallback
