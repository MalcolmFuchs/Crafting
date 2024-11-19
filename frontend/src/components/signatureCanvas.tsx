import { useEffect, useRef, useState } from "react";
import SignaturePad from "signature_pad";


interface SignaturePadProps {
  onSave: (dataURL: string) => void;
  onClear?: () => void;
  onCancel: () => void;
}

const SignaturePadComponent = ({ onSave, onClear, onCancel }: SignaturePadProps) => {
  const canvasRef                     = useRef<HTMLCanvasElement | null>(null);
  const signPadRef                    = useRef<SignaturePad | null>(null);


  const resizeCanvas = () => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const ratio   = Math.max(window.devicePixelRatio || 1, 1);
    canvas.width  = canvas.offsetWidth * ratio;
    canvas.height = canvas.offsetHeight * ratio;

    const context = canvas.getContext('2d');
    if (context) {
      context.scale(ratio, ratio);
    }

    if (signPadRef.current) {
      signPadRef.current.clear();
    }
  };


  useEffect(() => {
    const canvas = canvasRef.current;
    if (canvas) {
      signPadRef.current = new SignaturePad(canvas, {
        penColor: 'black',
      });
      resizeCanvas(); 
    }

    window.addEventListener('resize', resizeCanvas); 
    return () => {
      window.removeEventListener('resize', resizeCanvas);
    };
  }, []);

  const handleSave = () => {
    if (signPadRef.current) {
      const dataURL = signPadRef.current.toDataURL();
      onSave(dataURL);
    }
  };

  const handleClear = () => {
    if (signPadRef.current) {
      signPadRef.current.clear();
      if (onClear) onClear();
    }
  };

  return (
    <div className="relative flex flex-col items-center">
      <canvas
        ref={canvasRef}
        className="absolute top-0 border border-gray-300 bg-white w-full h-40"
      />
      <div className="flex space-x-4 mt-4">
      <button
          onClick={handleClear}
          className="px-4 py-2 bg-yellow-500 hover:bg-yellow-700 text-white rounded"
        >
          Abbrechen
        </button>
        <button
          onClick={onCancel}
          className="px-4 py-2 bg-yellow-500 hover:bg-yellow-700 text-white rounded"
        >
          Zurücksetzen
        </button>
        <button
          onClick={handleSave}
          className="px-4 py-2 bg-blue-500 hover:bg-blue-700 text-white rounded"
        >
          Speichern
        </button>
      </div>
    </div>
  );
};

export default SignaturePadComponent;