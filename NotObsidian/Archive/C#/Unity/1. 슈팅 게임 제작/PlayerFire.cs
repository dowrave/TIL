using System.Collections;
using System.Collections.Generic;
using JetBrains.Annotations;
using UnityEngine;

public class PlayerFire : MonoBehaviour
{
    public GameObject bulletFactory;

    // �Ѿ� ����
    public int poolSize = 10;

    // �迭 -> ����Ʈ�� ������Ʈ Ǯ ����
    public List<GameObject> bulletObjectPool;

    public GameObject firePosition;

    // �Ѿ� ������Ʈ Ǯ �����
    // 1. źâ�� ������ �� �����
    void Start()
    {
        // 2. ũ��� �Ѿ��� ���� ��ŭ
        bulletObjectPool = new List<GameObject>();
        // 3. źâ�� ���� ������ŭ �ݺ�
        for (int i = 0; i < poolSize; i++)
        {
            // 4. �Ѿ� ����
            GameObject bullet = Instantiate(bulletFactory);

            // 5. �Ѿ��� ������Ʈ Ǯ�� �ִ´�
            bulletObjectPool.Add(bullet);

            // ��Ȱ��ȭ
            bullet.SetActive(false);

        }

        #if UNITY_ANDROID
                GameObject.Find("Joystick canvas XYBZ").SetActive(true);
        #elif UNITY_EDITOR || UNITY_STANDALONE
                GameObject.Find("Joystick canvas XYBZ").SetActive(false);
        #endif

    }

    // Update is called once per frame
    void Update()
    {
#if UNITY_EDITOR || UNITY_STANDALONE
        if (Input.GetButtonDown("Fire1"))
        {
            Fire();
        }
#endif
    }

    public void Fire()
    { 
        if (bulletObjectPool.Count > 0)
        {
            GameObject bullet = bulletObjectPool[0];
            bullet.SetActive(true);
            bulletObjectPool.Remove(bullet);
            // �Ѿ� ��ġ��Ű��
            bullet.transform.position = transform.position;
        }
    }
}
