// src/components/RibbonItem.tsx
import React, { useState, useRef, useEffect } from 'react';

interface RibbonItemProps {
  label: string;
  children?: React.ReactNode;
}

const RibbonItem: React.FC<RibbonItemProps> = ({ label, children }) => {
  const [isOpen, setIsOpen] = useState(false);
  const itemRef = useRef<HTMLLIElement>(null);

  const handleToggle = () => {
    setIsOpen(!isOpen);
  };

  const handleClickOutside = (event: MouseEvent) => {
    if (itemRef.current && !itemRef.current.contains(event.target as Node)) {
      setIsOpen(false);
    }
  };

  useEffect(() => {
    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <li className="ribbon-item" ref={itemRef} onClick={handleToggle}>
      <span>{label}</span>
      {isOpen && (
        <ul className="ribbon-dropdown">
          {children}
        </ul>
      )}
    </li>
  );
};

export default RibbonItem;