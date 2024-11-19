"use client"

import SignaturePadComponent from "@/components/signatureCanvas";
import { format } from "path";
import { useEffect, useState } from "react"

export default function Page() {
  const [ isRunning, setIsRunning ]         = useState(false);
  const [ time, setTime ]                   = useState(0);
  const [ startTime, setStartTime ]         = useState<Date | null>(null);
  const [ endTime, setEndTime ]             = useState<Date | null>(null);
  const [ showSignField, setShowSignField ] = useState(false);

  let timer: NodeJS.Timeout | null = null;

  const handleStartStop = () => {
    if (isRunning) {
      setIsRunning(false);
      setShowSignField(true);
    } else {
      setTime(0);
      setIsRunning(true);
      setStartTime(new Date());
      setEndTime(null)
    }
  }

  const handleSaveSign = (dataURL: string) => {
    setEndTime(new Date());
    setShowSignField(false);
    alert("Vorgang abgeschlossen")
  }


  const handleCancelSignature = () => {
    setShowSignField(false);
    setIsRunning(true)
  }

  useEffect(() => {

    if (isRunning) {
      timer = setInterval(() => {
        setTime((prevTime) => prevTime + 1);
      }, 1000);
    }
    return () => {
      if (timer) {
        clearInterval(timer);
      }
    };
  }, [isRunning]);

  const formatTime = (seconds: number) => {
    const mins = Math.floor(seconds / 60).toString().padStart(2, "0")
    const secs = (seconds % 60).toString().padStart(2, "0")
    
    return `${mins}:${secs}`
  }

  const formatDate = (date: Date | null) => {
    if(!date) return "";
    return date.toLocaleTimeString();
  }

  return ( 
    <div className="grid gap-4 p-4">
      <h1>Zeiterfassung</h1>
        <h2 >Startzeit: {formatDate(startTime)}</h2>
        <h2 >Laufende Zeit: {formatTime(time)}</h2>
        <h2 >Endzeit: {formatDate(endTime)}</h2>
      <div className="grid grid-cols-2 gap-4">
        <button 
          onClick={handleStartStop}
          className={`${ isRunning ? 'bg-red-500 hover:bg-red-700' : 'bg-green-500 hover:bg-green-700'} text-white`}
        >
          <h2>{isRunning ? "Stop" : "Start"}</h2>
        </button>
        {showSignField && (
        <SignaturePadComponent 
          onSave={handleSaveSign}
          onCancel={handleCancelSignature}
        />
      )}
      </div>
    </div>
  )
}