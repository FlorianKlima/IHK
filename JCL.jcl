//P2123KC3 JOB ('ANR=N06#N927','BEZ=Insert-Programm WAP1'),             JOB23732
// 'KC21BN1',MSGCLASS=T,NOTIFY=P2123,MSGLEVEL=(2,1),
//  SCHENV=DB2P                                                         00000200
/*JOBPARM ROOM=013                                                      00000300
//AMS1     EXEC PGM=IDCAMS                                              00000500
//SYSPRINT DD SYSOUT=*                                                  00000600
  DELETE FBDV0001.KC21BN3.PROT                                          00000700
  IF MAXCC = 8 THEN SET MAXCC = 0                                           
//BKOSET   SET JOB=P2123KC3                                               
//*-------------------------------------------------------------------* 00001001
//ISRTSTEP EXEC DB2BATCH,MBR=KC21BN3,ENV=LE,LIB=EFEUL           
//G.PAN$SQL  DD *                                                     
PLAN=KC21BN3                                                                  
//G.SYSPRINT DD SYSOUT=*                                              
//G.EINGABE  DD DSN=FBDV0001.KC21BN3.EINGABE,DISP=SHR           
//G.PROT     DD DSN=FBDV0001.KC21BN3.PROT,                        
//           DISP=(,CATLG),                                       
//           SPACE=(TRK,(10,10),RLSE)                       
//*-------------------------------------------------------------------*