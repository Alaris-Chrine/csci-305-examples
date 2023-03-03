(defun testIf( alist )
    (cond
    ((null alist) 0)
    ((atom alist) 1)
    ((list alist) 33)
    (t 55)
   )
)
(write(testIf 'a))
(terpri)
(write(testIf '() ))
(terpri)
(write(testIf '(4 55)) )
(terpri)
(write(testIf '4))
(terpri)
(write(testIf nil))
	

