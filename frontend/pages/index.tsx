import React from 'react'
import { Button } from "@/components/ui/button"

export default function HomePage() {
  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-50 to-indigo-50 flex flex-col items-center justify-center p-8">
      <div className="w-full max-w-3xl bg-white rounded-lg shadow-xl overflow-hidden">
        <div className="p-6 bg-gradient-to-r from-blue-600 to-indigo-600">
          <h1 className="text-white text-2xl font-bold">Next.js with Tailwind CSS and shadcn/ui</h1>
          <p className="text-blue-100 mt-2">A demo showcasing shadcn/ui components with Tailwind styling</p>
        </div>
        
        <div className="p-6">
          <h2 className="text-xl font-semibold text-gray-800 mb-6">Button Component Examples</h2>
          
          <div className="grid grid-cols-2 gap-4 md:grid-cols-3">
            <div className="flex flex-col items-center gap-2 p-4 rounded-md border border-gray-200">
              <Button>Default Button</Button>
              <span className="text-sm text-gray-500">Default</span>
            </div>
            
            <div className="flex flex-col items-center gap-2 p-4 rounded-md border border-gray-200">
              <Button variant="secondary">Secondary</Button>
              <span className="text-sm text-gray-500">Secondary</span>
            </div>
            
            <div className="flex flex-col items-center gap-2 p-4 rounded-md border border-gray-200">
              <Button variant="destructive">Delete</Button>
              <span className="text-sm text-gray-500">Destructive</span>
            </div>
            
            <div className="flex flex-col items-center gap-2 p-4 rounded-md border border-gray-200">
              <Button variant="outline">Outline</Button>
              <span className="text-sm text-gray-500">Outline</span>
            </div>
            
            <div className="flex flex-col items-center gap-2 p-4 rounded-md border border-gray-200">
              <Button variant="ghost">Ghost</Button>
              <span className="text-sm text-gray-500">Ghost</span>
            </div>
            
            <div className="flex flex-col items-center gap-2 p-4 rounded-md border border-gray-200">
              <Button variant="link">Link Style</Button>
              <span className="text-sm text-gray-500">Link</span>
            </div>
          </div>
          
          <div className="mt-8 space-y-4">
            <h3 className="text-lg font-medium text-gray-800">Button Sizes</h3>
            <div className="flex flex-wrap items-center gap-4">
              <Button size="sm">Small</Button>
              <Button>Default</Button>
              <Button size="lg">Large</Button>
            </div>
          </div>
          
          <div className="mt-8 p-4 bg-gray-50 rounded-md">
            <p className="text-sm text-gray-600">
              This page demonstrates the successful integration of Next.js, Tailwind CSS, and
              the shadcn/ui component library. The buttons above are using shadcn/ui styling 
              with Tailwind utility classes.
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}