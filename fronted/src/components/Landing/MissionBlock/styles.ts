import styled from 'styled-components'

export const MissionBlockContainer = styled.section`
  position: relative;
  padding: 10rem 0 8rem;

  @media only screen and (max-width: 768px) {
    padding: 0rem 0 6rem;
  }
`

export const Content = styled.p`
  margin: 1.5rem 0 2rem 0;
`

export const ContentWrapper = styled.div`
  position: relative;
  max-width: 540px;

  @media only screen and (max-width: 480px) {
    margin: 2rem 0;
  }
`