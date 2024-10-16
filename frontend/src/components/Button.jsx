import React, { Children } from 'react'

const Button = ({bg, title, hover, margin, onClick}) => {
  return (
    <div>
      <button onClick={onClick} type='button' className={`bg-blue-500 hover:${hover} text-white font-bold py-2 px-4 rounded ${bg} ${margin}`}>{title}</button>
    </div>
  )
}

export default Button
